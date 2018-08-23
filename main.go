package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"strings"
	"net/url"
	"os"
	"path/filepath"
	"encoding/json"
	"github.com/go-kit/kit/log"
	"github.com/boltdb/bolt"
)

func main(){
	path := "ip/f.txt"
	path1 := "ip/tt.txt"
	path2 := "ip/iplist.txt"
	//pathcity := "getIp/ip/city.txt"
	//pathcitylist := "getIp/ip/citylist.txt"
	pathjson := "ip/provinces.json"

	logger := log.NewLogfmtLogger(os.Stdout)
	pathbolt := "ip/mydb.db"
	db, err := Connect(pathbolt)
	if err != nil {
		logger.Log("open",err)
	}
	defer db.Close()

	///redis
	//defer c.Close()
	//
	//_, err = c.Do("SET", "mykey", "superWang")
	//if err != nil {
	//	fmt.Println("redis set failed:", err)
	//}



	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		panic(err)
	}

	fi, err := os.Create(path1)
	if err != nil {
		panic(err)
	}
	defer fi.Close()

	f2, err := os.Create(path2)
	if err != nil {
		panic(err)
	}
	defer f2.Close()

	var result int64
	filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		result = f.Size()
		return nil
	})
	//fmt.Println(result)
	file := make([]byte,result)
	n1, err := f.Read(file)
	if err != nil{
		fmt.Println(err)
		fmt.Println(n1)
	}
	//t, _ := ioutil.ReadAll(f)
	str := string(file)
	//fmt.Println(n1)

	//************获取省市区名称编号****************
	f3, err := os.Open(pathjson)
	defer f.Close()
	if err != nil {
		panic(err)
	}
	t, _ := ioutil.ReadAll(f3)
	ss := string(t)
	m := make(map[string]map[string]string)
	err = json. Unmarshal ([]byte(ss), &m)
	if err != nil {
		fmt. Println ( "error:" , err )
	}
	//fmt. Printf ( "%+v" , m )
	//fmt. Println()
	//fmt. Println(m["86"])
	//***********************************************


	//var mm map[string]string
	mm := make(map[string]string)
	var s string
	var i int
	//maps := make(map[string]string)
	for _, j := range str{

		if j != '\n'{

			s += string(j)

		}else{
			i += 1
			//arr = append(arr, s)
			arri := strings.Fields(s)
			//fmt.Println(arri)
			//map1 := map[string]string{arri[0]: `{"ip":"` + arri[0] +`","address":"` + arri[2] +`","company":"` + arri[3] +`"}`}
			//maps = append(maps, map1) 

			//maps[arri[1]] = `{"ip":"` + arri[1] +`","address":"` + arri[2] +`","company":"` + arri[3] +`"}`
			//fmt.Println(len(arri))
			s1 := ""
			s2 := ""
			province := ""
			city := ""
			district := ""
			province_id := ""
			city_id := ""
			district_id := ""

			//正常
			for k, v := range m["86"]{
				if strings.Contains(arri[2], v){
					province = v
					province_id = k
					for kk, vv := range m[k]{
						if strings.Contains(arri[2], vv){
							city = vv
							city_id = kk
							for kkk, vvv := range m[kk]{
								if strings.Contains(arri[2], vvv){
									district = vvv
									district_id = kkk
								}
							}
						}
					}
				}
			}

			//内蒙古  "150000": "内蒙古自治区"
			if strings.Contains(arri[2], "内蒙古"){
				province = "内蒙古自治区"
				province_id = "150000"
				for kk, vv := range m["150000"]{
					if strings.Contains(arri[2], vv){
						city = vv
						city_id = kk
						for kkk, vvv := range m[kk]{
							if strings.Contains(arri[2], vvv){
								district = vvv
								district_id = kkk
							}
						}
					}
				}
			}

			//"450000": "广西壮族自治区",
			if strings.Contains(arri[2], "广西"){
				province = "广西壮族自治区"
				province_id = "450000"
				for kk, vv := range m["450000"]{
					if strings.Contains(arri[2], vv){
						city = vv
						city_id = kk
						for kkk, vvv := range m[kk]{
							if strings.Contains(arri[2], vvv){
								district = vvv
								district_id = kkk
							}
						}
					}
				}
			}

			//"540000": "西藏自治区",
			if strings.Contains(arri[2], "西藏"){
				province = "西藏自治区"
				province_id = "540000"
				for kk, vv := range m["540000"]{
					if strings.Contains(arri[2], vv){
						city = vv
						city_id = kk
						for kkk, vvv := range m[kk]{
							if strings.Contains(arri[2], vvv){
								district = vvv
								district_id = kkk
							}
						}
					}
				}
			}

			//"640000": "宁夏回族自治区",
			if strings.Contains(arri[2], "宁夏"){
				province = "宁夏回族自治区"
				province_id = "640000"
				for kk, vv := range m["640000"]{
					if strings.Contains(arri[2], vv){
						city = vv
						city_id = kk
						for kkk, vvv := range m[kk]{
							if strings.Contains(arri[2], vvv){
								district = vvv
								district_id = kkk
							}
						}
					}
				}
			}

			//"650000": "新疆维吾尔自治区"
			if strings.Contains(arri[2], "新疆"){
				province = "新疆维吾尔自治区"
				province_id = "650000"
				for kk, vv := range m["650000"]{
					if strings.Contains(arri[2], vv){
						city = vv
						city_id = kk
						for kkk, vvv := range m[kk]{
							if strings.Contains(arri[2], vvv){
								district = vvv
								district_id = kkk
							}
						}
					}
				}
			}

			//"810000": "香港特别行政区",
			if strings.Contains(arri[2], "香港"){
				province = "香港特别行政区"
				province_id = "810000"
				for kk, vv := range m["810000"]{
					if strings.Contains(arri[2], vv){
						city = vv
						city_id = kk
						for kkk, vvv := range m[kk]{
							if strings.Contains(arri[2], vvv){
								district = vvv
								district_id = kkk
							}
						}
					}
				}
			}

			//"820000": "澳门特别行政区"
			if strings.Contains(arri[2], "澳门"){
				province = "澳门特别行政区"
				province_id = "820000"
				for kk, vv := range m["820000"]{
					if strings.Contains(arri[2], vv){
						city = vv
						city_id = kk
						for kkk, vvv := range m[kk]{
							if strings.Contains(arri[2], vvv){
								district = vvv
								district_id = kkk
							}
						}
					}
				}
			}

			if len(arri)>3{
				if arri[3]=="联通" || arri[3]=="移动" || arri[3]=="电信"{
					s1 = `{"ip":"` + arri[0] +`","address":"` + arri[2] +`","company":"` + arri[3] +`","isp":"` + arri[3] +`","province":"` + province +`","city":"` + city +`","district":"` + district +`","province_id":"` + province_id +`","city_id":"` + city_id +`","district_id":"` + district_id +`"}`
					s2 = `{"ip":"` + arri[1] +`","address":"` + arri[2] +`","company":"` + arri[3] +`","isp":"` + arri[3] +`","province":"` + province +`","city":"` + city +`","district":"` + district +`","province_id":"` + province_id +`","city_id":"` + city_id +`","district_id":"` + district_id +`"}`
					province = ""
					city = ""
					district = ""
					province_id = ""
					city_id = ""
					district_id = ""
				}else {
					s1 = `{"ip":"` + arri[0] +`","address":"` + arri[2] +`","company":"` + arri[3] +`","isp":"","province":"` + province +`","city":"` + city +`","district":"` + district +`","province_id":"` + province_id +`","city_id":"` + city_id +`","district_id":"` + district_id +`"}`
					s2 = `{"ip":"` + arri[1] +`","address":"` + arri[2] +`","company":"` + arri[3] +`","isp":"","province":"` + province +`","city":"` + city +`","district":"` + district +`","province_id":"` + province_id +`","city_id":"` + city_id +`","district_id":"` + district_id +`"}`
					province = ""
					city = ""
					district = ""
					province_id = ""
					city_id = ""
					district_id = ""
				}
			}else{
				s1 = `{"ip":"` + arri[0] +`","address":"` + arri[2] +`","company":"","isp":"","province":"` + province +`","city":"` + city +`","district":"` + district +`","province_id":"` + province_id +`","city_id":"` + city_id +`","district_id":"` + district_id +`"}`
				s2 = `{"ip":"` + arri[1] +`","address":"` + arri[2] +`","company":"","isp":"","province":"` + province +`","city":"` + city +`","district":"` + district +`","province_id":"` + province_id +`","city_id":"` + city_id +`","district_id":"` + district_id +`"}`
				province = ""
				city = ""
				district = ""
				province_id = ""
				city_id = ""
				district_id = ""
			}

			//Save(db, arri[0], s1)
			mm[arri[0]] = s1
			//err = b.Put([]byte(arri[0]), []byte(s1))
			//w := s1 + "\n"
			//fi.WriteString(w)

			//f2.WriteString(arri[0] + "\n")


			if s1 != s2{
				//Save(db, arri[1], s1)
				//err = b.Put([]byte(arri[1]), []byte(s2))
				//w = s2 + "\n"
				mm[arri[1]] = s1
				//fi.WriteString(w) 
				//fmt.Println(w)
				//f2.WriteString(arri[1] + "\n") 
			}

			s=""

			if i%5000 == 0 || i >= 469636{
				fmt.Println(i)
				fmt.Println("*************************")

				Save(db, mm)
				mm = make(map[string]string)
			}
		}

	}

}

func Connect(path string)(*bolt.DB, error){
	db, err := bolt.Open(path, 0600, nil)

	return db, err
}
func Save(db *bolt.DB, mm map[string]string)error{

	err := db.Update(func(tx *bolt.Tx) error{
		//b := tx.Bucket([]byte("test"))
		b, err := tx.CreateBucketIfNotExists([]byte("test"))
		//dd := 1
		for key, val := range mm{
			//dd++
			//fmt.Println(dd)
			err = b.Put([]byte(key), []byte(val))
		}

		return err

	})
	return err
}

//func ConnectRedis(){
//	c, err := redis.Dial("tcp", "127.0.0.1:6379")
//	if err != nil {
//		fmt.Println("Connect to redis error", err)
//		return
//	}
//	
//}
func Get(key string)(string, error){
	logger := log.NewLogfmtLogger(os.Stdout)
	db, err := bolt.Open("ip/mydb.db", 0600, nil)
	if err != nil {
		logger.Log("open",err)
	}
	defer db.Close()
	var val string
	err = db.Update(func(tx *bolt.Tx)  error{
		b := tx.Bucket([]byte("test"))
		val= string(b.Get([]byte(key)))
		return err

	})
	return val, err
}


func httpP(){
	url1 := "http://ip.taobao.com/service/getIpInfo2.php"
	resp, err := http.PostForm(url1,
		url.Values{"ip": {"1.0.7.255"}})
	if err != nil {
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}
func httpDo() {

	client := &http.Client{}
	req, err := http.NewRequest("POST", "http://ip.taobao.com/service/getIpInfo2.php", strings.NewReader("ip=1.0.7.255"))
	if err != nil {
		// handle error 
	}
	cookie :="thw=cn; t=382ec866418922aa7fd24f86e901da98; cna=7Qr4E1ib4TgCAd3fU6aFL7qu; hng=CN%7Czh-CN%7CCNY%7C156; um=85957DF9A4B3B3E807D23256F3C852993D908631153C8E780BA9FD62CC4370BBBA2C98BB677CF5C7CD43AD3E795C914C6060EA8C3F0945E060A5885828FE2820; uc3=vt3=F8dBzrpKCaX1TAXXvAc%3D&id2=UoH4F%2BESjoXqJg%3D%3D&nk2=F5RHpxw9xa9YWVd3&lg2=UtASsssmOIJ0bQ%3D%3D; tracknick=tb2687612_11; lgc=tb2687612_11; _cc_=UtASsssmfA%3D%3D; tg=0; mt=ci=-1_0; x=e%3D1%26p%3D*%26s%3D0%26c%3D0%26f%3D0%26g%3D0%26t%3D0; cookie2=1f9342d5f10ce0e518560e8a49b0e1c8; v=0; _tb_token_=34ee9a3333e5a; uc1=cookie14=UoTfL8nEwpcMqA%3D%3D&lng=zh_CN&cookie16=UIHiLt3xCS3yM2h4eKHS9lpEOw%3D%3D&existShop=false&cookie21=UtASsssmfaCONGki4KTH3w%3D%3D&tag=8&cookie15=V32FPkk%2Fw0dUvg%3D%3D&pas=0; skt=729bfc5b6d8e4810; csg=5e2b5631; existShop=MTUzNDI5NzU2OQ%3D%3D; dnk=tb2687612_11; miid=48322721650242820; hibext_instdsigdipv2=1"
	req.Header.Set("Content-Type", "text/html;charset=UTF-8")
	req.Header.Set("Cookie", cookie)

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {

		// handle error

	}



	fmt.Println(string(body))

}
