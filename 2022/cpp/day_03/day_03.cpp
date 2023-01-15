#include <iostream>
#include <cstring>
#include <fstream>

#define LOG_D(fmt, ...) printf("%s[%d]: " fmt "\n", __FUNCTION__, __LINE__, ##__VA_ARGS__);
#define LOG_I(fmt, ...) printf(fmt "\n", ##__VA_ARGS__);

// #define PUZZLE_INPUT "day_03_sample.input.txt"
#define PUZZLE_INPUT "day_03.input.txt"

char find_common(std::string h1, std::string h2) {
    // check if common char is found in two strings
    for(int i = 0; i < (int)h1.length(); i++) {
        for(int j = 0; j < (int)h2.length(); j++) {
            if(h1[i] == h2[j]) {                       
                return h1[i];
            }
        }
    }
    return '\0';
}

char find_3common(std::string l1, std::string l2, std::string l3) {
    // check if common is found in three strings and return the common char 
    for(int i = 0; i < (int)l1.length(); i++) {
        for(int j = 0; j < (int)l2.length(); j++) {
            for(int k = 0; k < (int)l3.length(); k++) {
                if(l1[i] == l2[j] && l2[j] == l3[k]) {
                    return l1[i];
                }
            }
        }
    }
    return '\0';
}

int main() {

    std::ifstream ifs(PUZZLE_INPUT);
    std::string line;
    int pi = 0;
    int line_len = 0;
    char cchar;
    std::string h1, h2;
    int total_pri = 0;
    while(std::getline(ifs, line) || pi) {
        if(!line.empty()) {
            pi = 1;    
            line_len = line.length();
            h1 = line.substr(0, line_len/2);
            h2 = line.substr(line_len/2, line_len);
            cchar = find_common(h1, h2);
            // logic to calculate the priority and get the total
            if(cchar != '\0') {
                if((cchar - 96) > 1) {
                    total_pri = total_pri + (cchar - 96);
                } else {
                    total_pri = total_pri + (cchar - 38);
                }
            } else {    
                LOG_D("common item not found");
            }
        } else {
            pi = 0;
        }
    }
    LOG_I("total priority: %d", total_pri);
    int t3p = 0;    
    std::string l1, l2, l3;
    std::ifstream ifs3(PUZZLE_INPUT);
    pi = 0;
    while(std::getline(ifs3, l1) || pi) {
        if(!l1.empty()) {
            pi = 1;
            std::getline(ifs3, l2);
            std::getline(ifs3, l3);
            char c3char = find_3common(l1, l2, l3);
            // logic to calculate the group of three priority
            if(c3char != '\0') {
                if((c3char - 96) > 1) {
                    t3p = t3p + (c3char - 96);
                } else {
                    t3p = t3p + (c3char - 38);
                }
            }
        } else {
            pi = 0;
        }
    }
    LOG_I("group of three priority: %d", t3p);
    return 0;
}
