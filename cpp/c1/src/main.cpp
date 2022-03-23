
#include <iostream>
#include <fstream>
#include <time.h>
#include <string>
#include "nlohmann/json.hpp"

#include <sys/types.h>
#include <sys/socket.h>
#include <netinet/in.h>
#include <arpa/inet.h>


using namespace std;
// using json = nlohmann::json;
using json = nlohmann::json;

void ping() {
    cout << "pong" << endl;
}

void LoadJsonConfig()
{
    unsigned long destAddr;
    unsigned short destPort;
    unsigned int simID;
    map<string, string> scenarioMappings;
    map<string, int> scenarioIndices;
	int expectedMappings = 0;
	std::ifstream infile;
	json config;

    infile.open("data.json");
    try{
        config = nlohmann::json::parse(infile);
    } catch (json::exception& ex) {
        // asLog << "Failed to parse ScenarioMonitor.cfg: " << ex.what << endl;
        cout << "Failed to parse file: " << ex.what() << endl;
    	infile.close();
        return;
    }
	infile.close();

	if (config["SeraServerIp"] != nullptr) {
        string strAddr = config["SeraServerIp"];
		destAddr = inet_addr(strAddr.c_str());
	}
	if (config["SeraServerPort"] != nullptr) {
		destPort = config["SeraServerPort"];
	}
	if (config["SimId"] != nullptr) {
		simID = config["SimId"];
	}
	if (config["Mapped"] != nullptr) {
		expectedMappings = config["Mapped"];
	}
    if (expectedMappings <= 0) {
        return;
    }
    if (expectedMappings > 255) {
        expectedMappings = 255;
    }
	json data = config["Scenarios"];
	if (data != nullptr) {
        int countedMaps = 0;
		for (json obj : data) {
			string key;
			string value;
			if (obj["PreparedScenario"] != nullptr) {
				value = obj["PreparedScenario"];
			}
			if (obj["SeraScenario"] != nullptr) {
				key = obj["SeraScenario"];
				// key.resize(128, ' ');
			}
            countedMaps++;
            if (countedMaps > expectedMappings) {
                break;
            }
			scenarioMappings[key] = value;
		}
	}

    cout << "destAddr: " << destAddr << endl;
    cout << "destPort: " << destPort << endl;
    cout << "mappings: " << expectedMappings << endl;
    for (auto m : scenarioMappings) {
        cout << m.first << " --> " << m.second << endl;
    }
}

void foo() {
    cout << "Hello World!" << endl;
    ifstream i("./data.json");
    nlohmann::json j;
    // cout << "Start [" << i.rdbuf() << "] Done" << endl;
    i >> j;

    unsigned short destPort = 0;
    auto tmp = j["FakeData"]["number"];
    if (tmp != nullptr) {
        destPort = tmp;
    }

    cout << destPort << endl;

}

int main() {
    LoadJsonConfig();
    return 1;
}