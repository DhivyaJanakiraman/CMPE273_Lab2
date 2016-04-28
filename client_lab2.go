package main

import (
	//"ConsistHashing"
	"crypto/md5"
	//"encoding/json"
	"fmt"
	"io/ioutil"
	//"math"
	//"crypto/md5"
	"log"
	//"math"
	"net/http"
	"os"
	"sort"
	//"strconv"
	"strings"
)

type HashKey uint32
type HashKeyOrder []HashKey

func (h HashKeyOrder) Len() int           { return len(h) }
func (h HashKeyOrder) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h HashKeyOrder) Less(i, j int) bool { return h[i] < h[j] }

type HashRing struct {
	ring       map[HashKey]string
	sortedKeys []HashKey
	nodes      []string
}

func New(nodes []string) *HashRing {
	hashRing := &HashRing{
		ring:       make(map[HashKey]string),
		sortedKeys: make([]HashKey, 0),
		nodes:      nodes,
	}
	hashRing.generateCircle()
	return hashRing
}

func (h *HashRing) generateCircle() {
	j := 0
	for _, node := range h.nodes {
		nodeKey := fmt.Sprintf("%s-%d", node, j)
		bKey := HashMD5(nodeKey)
		j++
		for i := 0; i < 3; i++ {
			key := hashVal(bKey[i*4 : i*4+4])
			h.ring[key] = node
			h.sortedKeys = append(h.sortedKeys, key)
		}
	}

	sort.Sort(HashKeyOrder(h.sortedKeys))
}

func (h *HashRing) GetNode(stringKey string) (node string, ok bool) {
	pos, ok := h.GetNodePos(stringKey)
	if !ok {
		return "", false
	}
	return h.ring[h.sortedKeys[pos]], true
}

func (h *HashRing) GetNodePos(stringKey string) (pos int, ok bool) {
	if len(h.ring) == 0 {
		return 0, false
	}

	key := h.GenKey(stringKey)

	nodes := h.sortedKeys
	pos = sort.Search(len(nodes), func(i int) bool { return nodes[i] > key })

	if pos == len(nodes) {
		// Wrap the search, should return first node
		return 0, true
	} else {
		return pos, true
	}
}

func (h *HashRing) GenKey(key string) HashKey {
	bKey := HashMD5(key)
	return hashVal(bKey[0:4])
}

func (h *HashRing) AddNode(node string) *HashRing {
	nodes := make([]string, len(h.nodes), len(h.nodes)+1)
	copy(nodes, h.nodes)
	nodes = append(nodes, node)

	hashRing := &HashRing{
		ring:       make(map[HashKey]string),
		sortedKeys: make([]HashKey, 0),
		nodes:      nodes,
	}
	hashRing.generateCircle()
	return hashRing

}

func (h *HashRing) RemoveNode(node string) *HashRing {
	nodes := make([]string, 0)
	for _, eNode := range h.nodes {
		if eNode != node {
			nodes = append(nodes, eNode)
		}
	}

	hashRing := &HashRing{
		ring:       make(map[HashKey]string),
		sortedKeys: make([]HashKey, 0),
		nodes:      nodes,
	}
	hashRing.generateCircle()
	return hashRing
}

func hashVal(bKey []byte) HashKey {
	return ((HashKey(bKey[3]) << 24) |
		(HashKey(bKey[2]) << 16) |
		(HashKey(bKey[1]) << 8) |
		(HashKey(bKey[0])))
}

func HashMD5(key string) []byte {
	m := md5.New()
	m.Write([]byte(key))
	return m.Sum(nil)
}

var Ring *HashRing
var myServers []string
var myData = make(map[string]string)

//Main function for Client, if Client is to be run as a Simple Standalone GO application
func main() {
	myServers = []string{
		"127.0.0.1:3001",
		"127.0.0.1:3002",
		"127.0.0.1:3003",
		"127.0.0.1:3004",
		"127.0.0.1:3005"}

	Ring = New(myServers)
	argument2 := os.Args[2]
	fmt.Printf("Argument2: %s", argument2)

	//Generate seed-data to be distributed across the servers
	GenMappings(argument2)

	//Distribute keys
	for key, value := range myData {
		fmt.Println("Distributing/Saving this <key-valye> pair : <", key, "-", value, ">")
		PutOperation(key, value)
	}

	//Try to retrive few keys
	fmt.Println("Retrieving FEW keys now!")
	GetOperation("2")
	GetOperation("4")
	GetOperation("1")

	//Try to retrive few keys
	fmt.Println("Retrieving ALL keys now!")

	GetAllOperation("127.0.0.1:3001")
	GetAllOperation("127.0.0.1:3002")
	GetAllOperation("127.0.0.1:3003")
	GetAllOperation("127.0.0.1:3004")
	GetAllOperation("127.0.0.1:3005")

}

func GenMappings(values string) {

	f := func(c rune) bool {
		return c == ',' || c == '-' || c == '>'
	}
	// Separate into fields with func.
	fields := strings.FieldsFunc(values, f)
	fmt.Println(fields)
	fmt.Println(len(fields))
	for i := 0; i < len(fields); i = i + 2 {

		myData[fields[i]] = fields[i+1]
		fmt.Printf("key : %s, value : %s\n", fields[i], myData[fields[i]])
	}
}

func PutOperation(key string, value string) {
	// Grab id
	id := key

	server, _ := Ring.GetNode(id)
	fmt.Println("The <key-value> pair to be saved : <", key, "-", value, ">")
	fmt.Println("The server for this key : ", server)

	//make a corresponding PUT request to the server here
	url := "http://" + server + "/keys/" + id + "/" + value

	client := &http.Client{}
	request, err := http.NewRequest("PUT", url, strings.NewReader(""))
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	} else {
		defer response.Body.Close()
		_, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("Key saved successfully!")
	fmt.Println("------------------------------------------------------------")
}

/*
* GetOperation function - Function to support GET operation and retrieve keys from sharded the data set into three server instances
 */
func GetOperation(key string) {
	// Grab id
	id := key

	server, _ := Ring.GetNode(id)
	fmt.Println("Retrieving key from this server : ", server) //127.0.0.1:3001

	//make a corresponding GET request to the server here
	url := "http://" + server + "/keys/" + id

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("The Response is : %s\n", string(contents))
	}

	fmt.Println("Key retrieved successfully!")
	fmt.Println("------------------------------------------------------------")

}

/*
* GetAllOperation function - Function to support GET ALL operation and retrieve ALL keys from any particular server instances
 */
func GetAllOperation(server string) {
	fmt.Println("Retrieving ALL keys from this server : ", server)

	//make a corresponding GET request to the server here
	url := "http://" + server + "/keys"

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("The Response is : %s\n", string(contents))
	}

	fmt.Println("Keys retrieved successfully!")
	fmt.Println("------------------------------------------------------------")

}
