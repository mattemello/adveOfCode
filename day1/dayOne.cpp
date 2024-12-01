#include <algorithm>
#include <charconv>
#include <cstdio>
#include <fstream>
#include <iostream>
#include <map>
#include <ostream>
#include <string>
#include <unordered_map>

int* split_string(std::string text) {

    int firstNum = 0;
    std::string* toReturn = new std::string[2];
    int* num = new int[2];

    int i = 0;

    while (text[i] != '\0') {

        if(text[i] == ' '){
            if(firstNum == 0) {
                firstNum++;
            }

        }else{
            toReturn[firstNum] += text[i]; 
        }

        i++;
    }


    num[0] = std::stoi(toReturn[0]);
    num[1] = std::stoi(toReturn[1]);


    return num;

}

int absoluteValue(int num) {

    if(num < 0) {
        return (num * -1);
    }

    return num;

}

int firstPart(int numLines, std::ifstream &readfile) {

    std::string text;
    int numberOne[numLines];
    int numberTwo[numLines];

    int pos1 = 0, pos2 = 0;

    int* returned;

    while (getline(readfile, text)) {

        returned = split_string(text);

        numberOne[pos1] = returned[0];
        numberTwo[pos2] = returned[1];

        pos1++;
        pos2++;
    }

    std::sort(numberOne, numberOne + (sizeof(numberOne) / sizeof(numberOne[0])));
    std::sort(numberTwo, numberTwo + (sizeof(numberTwo) / sizeof(numberTwo[0])));

    int sum = 0;

    for(int i = 0; i < numLines; i++) {

        sum += absoluteValue(numberOne[i] - numberTwo[i]); 

    }

    return sum;
}

int secondPart(int numLines, std::ifstream &readfile){
    std::string text;

    int pos2 = 0;
    int* returned;

    std::map<int, int> collonOne;
    std::map<int, int> repeat;
    int numberTwo[numLines];

    while (getline(readfile, text)) {

        returned = split_string(text);

        collonOne[returned[0]] = 0;
        if(collonOne.find(returned[0]) != collonOne.end()) {
            repeat[returned[0]] += 1;
        }

        numberTwo[pos2] = returned[1];

        pos2++;
    }

    for(int i = 0; i < pos2; i++) {

        if(collonOne.find(numberTwo[i]) != collonOne.end()){
            collonOne[numberTwo[i]] += 1;
        }

    }

    int molt = 0;

    for (auto value : collonOne) {

        if(repeat.find(value.first) != repeat.end()) {
            molt += ((value.first * value.second) * repeat[value.first]); 
        }else {
            molt += (value.first * value.second);
        }
    }
    
    return molt;
}

int main() {


    std::ifstream readfile("day1");

    std::string unused;

    int numLines = 0;

    while ( std::getline(readfile, unused) )
        ++numLines;

    readfile.clear();
    readfile.seekg(0);

    int sum = firstPart(numLines, readfile);

    readfile.clear();
    readfile.seekg(0);
    
    int count = secondPart(numLines, readfile);

    std::cout << sum << std::endl;
    std::cout << count << std::endl;
}
