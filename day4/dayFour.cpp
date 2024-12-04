#include <cstdio>
#include <cstdlib>
#include <iostream>
#include <fstream>
#include <ostream>
#include <string>

int controlXmas(char m, char a, char s){

    if(m == 'M'  && a == 'A' && s == 'S') {
        return 1;
    }

    return 0;

}

int firstPart(std::string* file, int dimension) {

    int countFile = 0;

    for (int i = 0; i < dimension-1; i++) {
        for(int j = 0; j < file[i].length()-1; j++) {

            if(file[i][j] == 'X'){
                if (i > 3) {
                    if(j > 3){
                        //top left
                        countFile += controlXmas(file[i-1][j-1],file[i-2][j-2],file[i-3][j-3]);
                    }
                    if(dimension - j > 3){
                        //top right
                        countFile += controlXmas(file[i-1][j+1],file[i-2][j+2],file[i-3][j+3]);
                    }
                    //top
                    countFile += controlXmas(file[i-1][j],file[i-2][j],file[i-3][j]);
                }
                if (dimension - i > 3){
                    if(j > 3){
                        //bottom left
                        countFile += controlXmas(file[i+1][j-1],file[i+2][j-2],file[i+3][j-3]);
                    }
                    if(dimension - j > 3){
                        //bottom right
                        countFile += controlXmas(file[i+1][j+1],file[i+2][j+2],file[i+3][j+3]);
                    }
                    //bottom
                    countFile += controlXmas(file[i+1][j],file[i+2][j],file[i+3][j]);
                }
                if(j > 3){
                    //left
                    countFile += controlXmas(file[i][j-1],file[i][j-2],file[i][j-3]);
                }
                if(dimension - j > 3){
                    //right
                    countFile += controlXmas(file[i][j+1],file[i][j+2],file[i][j+3]);
                }
            }

        }
    
    }

    return countFile;

}

int main() {
    
    std::string text;

    std::ifstream readFile("day4");
    int dimensionFile = 0;

    while (getline(readFile, text)) {
        dimensionFile++;
    }

    readFile.close();
    readFile.open("day4");

    std::string* allTheFile; 

    allTheFile = (std::string*) malloc(dimensionFile * sizeof(std::string*));

    int position = 0;

    while (getline(readFile, text)) {
        allTheFile[position] = text;
        std::cout << "dimension file: "<< position << std::endl;
        position++;
    }

    int response1 = firstPart(allTheFile, dimensionFile);
    
    std::cout << "responce 1: " << response1 << std::endl;
}
