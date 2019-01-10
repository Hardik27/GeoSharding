package clustering

import (
	"../gomeans"
	"flag"
	"math/rand"
	"os"
	"encoding/csv"
   "io"
   "fmt"
	"time"
	"strconv"
)
var count int=0
var n int=800
var port int=8000
var dataset []gomeans.Point


func GetClusters()  [][]int {
	rand.Seed(time.Now().UnixNano())
	k:=3
	// k := flag.Int("k", 0, "number of clusters")
	//size := flag.Int("n", 0, "number of elements")
	size:=800
	flag.Parse()

	if k == 0 || size == 0 {
		flag.Usage()
		os.Exit(1)
	}

	
	m,_:=parseLocation("data/dataset.csv")
	// fmt.Println("Printing Map")
	 fmt.Println(m["1.0.0.0/24"])
	fmt.Println(count) 
	var startTime=time.Now()
	fmt.Println("Starting clustering with drawing")
	clusters:=gomeans.RunWithDrawing(dataset, k)
	fmt.Println("Time elapsed")
	fmt.Println(time.Since(startTime))

	
	// for i:=0;i<len(clusters);i++{
	// 	//fmt.Println("Starting Writing to file")
	// 	//csvOut, _ := os.Create("Shard"+strconv.Itoa(i+1)+".csv")
	// 	//writer := csv.NewWriter(csvOut)
	// 	length:=len(clusters[i].Points)
	// 	fmt.Println("Length is: "+strconv.Itoa(length))
	// 	//sampling without replacements
	// 	fmt.Println("Starting Sampling")
	// 	indexes := make([]bool,length)	
	// 	for i := 0; i < n; i++ {
	// 		var r int
	// 		for {
	// 			r = rand.Intn(length)
	// 			fmt.Println(r)
	// 			if indexes[r] {
	// 				continue
	// 			}
	// 			break
	// 			}
	// 		indexes[r] = true
	// 	}


	// 	// //form csv for mapping on the world map
	// 	// for j:=0;j<len(clusters[i].Points);j++{
			
	// 	// 	if indexes[j]==true{
	// 	// 		var record []string
	// 	// 		record=append(record,strconv.Itoa(i+1))
	// 	// 		record=append(record,fmt.Sprintf("%f", clusters[i].Points[j].Lat))
	// 	// 		record=append(record,fmt.Sprintf("%f", clusters[i].Points[j].Lon))
	// 	// 		record=append(record,clusters[i].Points[j].Ip)
	// 	// 		writer.Write(record)
	// 	// 		record=nil
	// 	// 	}
			
	// 	// }
	// 	// indexes=nil
	// }
	fmt.Println("Done Clustering")
	finalMap:=GetRandomPoints(clusters)
	var temp [][]int
	fmt.Println("Final Map is")
	fmt.Println(finalMap)
	for i:=0;i <len(clusters);i++{
		var shardlist []int
		for j:=0;j<len(clusters[i].Points);j++{
			
			if val, ok := finalMap[clusters[i].Points[j].Ip]; ok {
				shardlist=append(shardlist,val)
			}	
			// shardlist=append(shardlist,finalMap[clusters[i].Points[j].Ip])
		}
		if len(shardlist)!=0{
			temp=append(temp,shardlist)
		}

		shardlist=nil
		fmt.Println("Temp now is: ")
		fmt.Println(temp)
	}
	return temp
}

func parseLocation(file string) (map[string]*gomeans.Point, error) {
	
	fmt.Println("reading dataset now")
	rFile, err := os.Open(file) //3 columns
  	if err != nil {
    	fmt.Println("Error:", err)
    	return nil,nil
   	}
  	defer rFile.Close()

    csvr := csv.NewReader(rFile)
	
    locations := map[string]*gomeans.Point{}
	fmt.Println("Made locations map")
	// 	indexes := make([]bool,length)	
	// 	for i := 0; i < n; i++ {
	// 		var r int
	// 		for {
	// 			r = rand.Intn(length)
	// 			fmt.Println(r)
	// 			if indexes[r] {
	// 				continue
	// 			}
	// 			break
	// 			}
	// 		indexes[r] = true
	// 	}
	
	for {
        row, err := csvr.Read()
        if err != nil {
            if err == io.EOF {
				fmt.Println("Done Reading")
				err = nil
            }
            return locations, err
		}
		// fmt.Println("No error for now")
		// fmt.Println(row)
		// fmt.Println(row[3]+" "+row[1]+" "+row[2])
		if row[1]!="" && row[2]!="" && row[3]!="" {     //1=7,2=8,3=0
			
			count++
        	p := &gomeans.Point{}
        	if p.Lat, err = strconv.ParseFloat(row[1], 64); err != nil {
				fmt.Println("error occurred1")
				fmt.Println(err)
            	return locations, err
        	}
        	if p.Lon, err = strconv.ParseFloat(row[2], 64); err != nil {
				fmt.Println("error occurred2")
				return locations, err
			}
			p.Ip=row[3]
			locations[row[3]] = p
			dataset = append(dataset, gomeans.Point{p.Lat,p.Lon,p.Ip})
		} else{
		}	
    }
}
func GetRandomPoints(clusters []gomeans.Cluster) map[string]int {



	rand.Seed(time.Now().UnixNano())
	fmt.Println(clusters[0].Points[10])
	fmt.Println(clusters[0].Points[20])
	fmt.Println(clusters[2].Points[10])
	fmt.Println(clusters[2].Points[20])
	m:=make(map[string]int)
	m[clusters[0].Points[10].Ip]=8000
	m[clusters[0].Points[20].Ip]=8001
	m[clusters[2].Points[10].Ip]=8002
	m[clusters[2].Points[20].Ip]=8003


	// index1:=rand.Intn(len(clusters))
	// fmt.Println("Random clsuter index")
	// fmt.Println(index1)
	
	// var index2 int
	
	// //choose 2 random clusters
	
	
	
	// for {
	// 	r:=rand.Intn(len(clusters))
	// 	if r==index1{
	// 		continue
	// 	}else{
	// 		index2=r
	// 		break
	// 	}
	// }
	// fmt.Println(index2)
	// fmt.Println("Getting random Points")
	// m:=make(map[string]int)
	
	
	
	
	// for i:=0;i<2;i++ {
		
	// 	for {
	// 		random:=rand.Intn(len(clusters[index1].Points))
	// 		if _, ok := m[clusters[index1].Points[random].Ip]; !ok {
				
	// 			m[clusters[index1].Points[random].Ip]=port
	// 			port++
	// 			break
	// 		}
	// 	}
	// }
	// fmt.Println("Getting next 2 nodes")
	// for i:=0;i<2;i++ {
		
	// 	for {
	// 		random:=rand.Intn(len(clusters[index2].Points))
	// 		if _, ok := m[clusters[index2].Points[random].Ip]; !ok {
	// 			m[clusters[index2].Points[random].Ip]=port
	// 			port++
	// 			break
	// 		}
	// 		// if m[clusters[index2].Points[random].Ip]==0{
	// 		// 	m[clusters[index2].Points[random].Ip]=port
	// 		// 	port++
	// 		// 	break
	// 		// }
	// 	}
		
	// }
	fmt.Println(m)
	return m
}