//back end for cloud service.
/**

1. mqtt subscriber 
2. takes in data from client
3. founds what io link master is being used.
    a. finds diagnostics of io link master.
    b. each topic will be different for each io link master.
    c. will have handle and support different io link masters. Starting with MUrr 54631 
4. finds the ports that are being used.
5. finds the sensors.
6. compiles useable data and sends to influxdb 

*/ 

package main

import (
    "fmt"
    "os" 
    "time"
    
    mqtt "github.com/eclipse/paho.mqtt.golang"
)


/** getIOLinkName
* 
* parse topic to find the iolink master's namne 
* Once found call the specific ioLink master's getDiagnostics/getSensorData 
*
*/

function getIOLinkName() String {

return ""
}



function main() {



}


