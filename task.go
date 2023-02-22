package main

import "fmt"

func main(){
fmt.Println(stringCmp("hello","holly"))
fmt.Println(stringReplace("Hi how are you","ar","hoy"))
///////////////////////////0123456789ABCD
}

func stringCmp(str1 string,str2 string)bool{
	s1:=[]rune(str1)
	s2:=[]rune(str2)
	if(stringLength(str1)!=stringLength(str2)){
		return false
	}
	n:=arrayLength(s1)
	for i:=0;i<n;i++{
		if(s1[i]!=s2[i]){
			return false
		}else{
			if(i==n-1){
				return true
			}
		}
	}
	return false

}
func arrayLength(arr []rune)int{
	count:=0
	for range arr{
		count+=1
	}
	return count
}
func stringLength(str string)int{
	s:=[]rune(str)
	return arrayLength(s)
}
func replaceFunc(sentence []rune,replace []rune,start int,end int)[]rune{
	length:=arrayLength((sentence))+arrayLength(replace)-(end-start)
	fmt.Println((length))
	fmt.Println(sentence)
	var returnArray = make([]rune,length)
	firstArray:=sentence[:start]
	fmt.Println(firstArray)
	secondArray:=sentence[end:]
	fmt.Println(secondArray)
	i:=0
	for ;i<=arrayLength(firstArray)-1;i++{
			returnArray[i]=firstArray[i]
	}
	fmt.Println(returnArray)
	for j:=0;j<=arrayLength(replace)-1;j++{
			returnArray[i]=replace[j]
			i+=1
	}
	fmt.Println(returnArray)
	for j:=0;j<arrayLength(secondArray);j++{
			returnArray[i]=secondArray[j]
			i+=1
	}
	fmt.Println(returnArray)
	return returnArray
}
func stringReplace(sentence string,replace string,new string)string{
	sentenceArray:=[]rune(sentence)
	replaceArray:=[]rune(replace)
	newArray:=[]rune(new)
	n:=arrayLength(replaceArray)
	sLen:=arrayLength(sentenceArray)
	start:=0
	isFound:=false
	for i:=0;i<arrayLength(sentenceArray) && !isFound;i++{
		
		if sentenceArray[i]==' '||i==sLen-1{
			fmt.Println(start,i-1)
			if i-start ==n {
				k:=0
				for j:=start;j<start+n;j++{
					fmt.Println(sentenceArray[j],replaceArray[k])
					if sentenceArray[j]!=replaceArray[k]{
						break
					}else{
						k+=1
						if j==i-1{
							fmt.Println("found")
							sentenceArray=replaceFunc(sentenceArray,newArray,start,start+n)
							isFound=true
						}
					}
				}
			}
			start = i+1
		}
	}
	return string(sentenceArray)
}

