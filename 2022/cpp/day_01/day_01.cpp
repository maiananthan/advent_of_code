#include <iostream>
#include <cstring>
#include <fstream>
#include <sstream>
#include <bits/stdc++.h>

#define LOG_D(fmt, ...) printf("%s[%d]: " fmt "\n", __FUNCTION__, __LINE__, ##__VA_ARGS__);
#define LOG_I(fmt, ...) printf(fmt "\n", ##__VA_ARGS__);

// puzzle input is saved to a file

int main() {

    int total_cal, max_cal, cal, pi;
    total_cal = max_cal = cal = pi = 0;
    int c[3] = {0, 0, 0};

    int num = sizeof(c) / sizeof(c[0]);
    
    // open the file to ifstream
    std::ifstream ifs("input.txt");
    if(ifs.is_open()) {
        std::string line;
        // get the line from the ifstream
        // for last group of entries, check if any previous inputs present
        while(std::getline(ifs, line) || pi) {
            // check if line is empty
            if(!line.empty()) {
                // make previous input is present
                pi = 1;
                // convert string to integer
                cal = std::stoi(line);
                // add the calories to total calories
                total_cal = total_cal + cal;
            } else {
                // to track the big three 
                // check if the last entry is smaller than current one
                // if so replace it
                if(c[2] < total_cal) {
                    c[2] = total_cal;
                }
                // sort the array from larger to smaller
                std::sort(c, c+num, std::greater<int>());
                
                // make the previous input flag, total_cal, cal to zero
                pi = total_cal = cal = 0;
            }
        }
    }
    LOG_I("top three: %d %d %d -> total: %d", c[0], c[1], c[2], c[0] + c[1] + c[2]);
    LOG_I("max_cal: %d", c[0]);
    return 0;
}
