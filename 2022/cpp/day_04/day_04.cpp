#include <iostream>
#include <fstream>
#include <cstring>

#define LOG_D(fmt, ...) printf("%s[%d]: " fmt "\n", __FUNCTION__, __LINE__, ##__VA_ARGS__);
#define LOG_I(fmt, ...) printf(fmt "\n", ##__VA_ARGS__);

int find_max(int i1, int i2, int i3, int i4) {
    int max = i1 > i2 ? i1 : i2;
    max = max > i3 ? max : i3;
    max = max > i4 ? max : i4;
    return max;
}

int overlap(int fs, int fe, int ss, int se) {
    int overlap = 0;
    if( (fs >= ss) && (fe <= se) ) {
        overlap = 1;
    } else if ( (ss >= fs) && (se <= fe) ) {
        overlap = 1;
    }
    return overlap;
}

int find_range_overlap(int fs, int fe, int ss, int se) {
    int range_overlap = 0;
    for(int i = fs; i <= fe; i++) {
        for(int j = ss; j <= se; j++) {
            if(i == j) {
                return 1;
            }
        }
    }
    return range_overlap;
}
int main() {
    std::ifstream ifs(PUZZLE_INPUT);
    
    std::string line;
    int pi = 0;
    std::string fst, snd;
    int fst_s, fst_e, snd_s, snd_e;
    int t_ol = 0;
    int t_r_ol = 0;
    while(std::getline(ifs, line) || pi) {
        if(!line.empty()) {
            pi = 1;
            fst = line.substr(0, line.find(","));
            snd = line.substr(line.find(",") + 1, line.length());
            LOG_D("fst: %s, snd: %s", fst.c_str(), snd.c_str());
            fst_s = std::stoi(fst.substr(0, fst.find("-")));
            fst_e = std::stoi(fst.substr(fst.find("-") + 1, fst.length()));
            snd_s = std::stoi(snd.substr(0, snd.find("-")));
            snd_e = std::stoi(snd.substr(snd.find("-") + 1, snd.length()));
            LOG_D("fst_s: %d, fst_e: %d, snd_s: %d, snd_e: %d", fst_s, fst_e, snd_s, snd_e);
            LOG_D("max: %d", find_max(fst_s, fst_e, snd_s, snd_e));
            t_ol += overlap(fst_s, fst_e, snd_s, snd_e);
            t_r_ol += find_range_overlap(fst_s, fst_e, snd_s, snd_e);
        } else {
            pi = 0;
        }
    }
    LOG_I("total overlap found: %d", t_ol);
    LOG_I("total range overlap found: %d", t_r_ol);
    return 0;
}
