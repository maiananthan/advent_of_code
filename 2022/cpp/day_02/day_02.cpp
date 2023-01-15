#include <iostream>
#include <sstream>
#include <fstream>
#include <cstring>

#define LOG_D(fmt, ...) printf("%s[%d]: " fmt "\n", __FUNCTION__, __LINE__, ##__VA_ARGS__);
#define LOG_I(fmt, ...) printf(fmt "\n", ##__VA_ARGS__);

int main() {
    int total_score = 0;
    std::ifstream ifs("input.txt");

    // std::stringstream ifs_ss;
    // ifs_ss << ifs.rdbuf();

    // std::string ifs_str = ifs_ss.str();

    // LOG_D("ifs_str: \'%s\'", ifs_str.c_str());
#if 0
    while(ifs >> a >> b) {
        LOG_D("a: %c, b: %c", a, b);
    }
#endif
    // if(ifs.is_open()) 
    char a;
    char b;
    int t2 = 0;
    int pi = 0;
    {
        std::string line;
        while(std::getline(ifs, line) || pi) {
            if(!line.empty()) {
                pi = 0;
            
                LOG_D("line: %s", line.c_str());
                a = line[0];
                b = line[2];
                // calculate your input
                //
                if(b == 'X') {
                    total_score += 1;
                } else if(b == 'Y') {
                    total_score += 2;
                } else if(b == 'Z') {
                    total_score += 3;
                }

                // check if you are won or lost or draw 
                // A B C 65 66 67
                // X Y Z 88 89 90
                // R P S
                if((a == 65 && b == 89) || (a == 66 && b == 90) || (a == 67 && b == 88)) {
                    // won
                    LOG_D("won");
                    total_score += 6;
                } else if((a == 65 && b == 90) || (a == 66 && b == 88) || (a == 67 && b == 89)) {
                    // lost
                    LOG_D("lost");
                    total_score += 0;
                } else if((a == 65 && b == 88) || (a == 66 && b == 89) || (a == 67 && b == 90)) {
                    // draw
                    LOG_D("draw");
                    total_score += 3;
                }
                
                if(b == 'X') {
                    t2 += 0;
                } else if(b == 'Y') {
                    t2 += 3;
                } else if(b == 'Z') {
                    t2 += 6;
                }
                if(b == 'X') {
                    // lose 
                    if(a == 65) {
                        t2 += 3; // 
                    } else if(a == 66) {
                        t2 += 1;
                    } else if(a == 67) {
                        t2 += 2;
                    }
                } else if(b == 'Y') {
                    // draw
                    if(a == 65) {
                        t2 += 1;
                    } else if(a == 66) {
                        t2 += 2;
                    } else if(a == 67) {
                        t2 += 3;
                    }
                } else if(b == 'Z') {
                    // win
                    if(a == 65) {
                        t2 += 2;
                    } else if(a == 66) {
                        t2 += 3;
                    } else if(a == 67) {
                        t2 += 1;
                    }
                }
                LOG_D("a: %c, b: %c, t2: %d", a, b, t2);
#if 0
                if((b == 'X') && (a == 'A')) {
                    // lose
                } else if()  {

                } else if() {

                }
#endif
                // LOG_D("a: %c, b: %c", a, b);
                // LOG_D("a: %d, b: %d", a, b);
            }

        }

        LOG_I("total score: %d, t2: %d", total_score, t2);
    }


    return 0;
}
