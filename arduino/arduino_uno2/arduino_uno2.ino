//WaterTemp Libs
#include <OneWire.h>
#include <DallasTemperature.h>
//Adafruit Humidity and Temp
#include <Wire.h>
#include <Adafruit_AM2315.h>

//pH sensor
#include <SoftwareSerial.h>

//RC Emitter
#include <RCSwitch.h>

//Define for Emitter
RCSwitch sender = RCSwitch();

// Led Pins for Modules
const int module1 = 13;
const int module2 = 11;
const int module3 = 10;
const int module4 = 9;
const int module5 = 8;
const int module6 = 7;


//sensor Pins
  //Legacy
const int tempPin = 0;
int photocellPin = 1;
  //UltraSonic
int trigPin = 5;
int echoPin = 6;
long duration;
  //WaterTemp
#define ONE_WIRE_BUS 4
OneWire oneWire(ONE_WIRE_BUS);
DallasTemperature sensors(&oneWire);
 //HumidityAndTemp
Adafruit_AM2315 am2315;
struct temphum {
  float temp;
  float humi;
};
 //pH
 #define rx 2
 #define tx 3
SoftwareSerial myserial(rx, tx);
float pH;
String sensorstring = "";
boolean sensor_string_complete;

  
// Dynamic Variables for LED blink
int ledPin;
bool blinkLed;

bool ledOn = true;

// Variables for Sensors
int printPeriod = 1000 * 10; // every Minute
unsigned long time_now = 0;


void setup() {
  //LightSetup
  pinMode(module1, OUTPUT);
  pinMode(module2, OUTPUT);
  pinMode(module3, OUTPUT);
  pinMode(module4, OUTPUT);
  pinMode(module5, OUTPUT);
  pinMode(module6, OUTPUT);

  digitalWrite(module1, HIGH);
  digitalWrite(module2, HIGH);
  digitalWrite(module3, HIGH);
  digitalWrite(module4, HIGH);
  digitalWrite(module5, HIGH);
  digitalWrite(module6, HIGH);

  //UltraSonic Setup
  pinMode(trigPin, OUTPUT);
  pinMode(echoPin, INPUT);

  //Emitter
  sender.enableTransmit(3);  // An Pin 3

  sender.setProtocol(1);
  sender.setPulseLength(300);



  //WaterTemp
  sensors.begin();

  //pH Node
  sensorstring.reserve(30);
  
  Serial.begin(9600);
  Serial3.begin(9600);
  myserial.begin(9600);


   //Wait for HumTemp to start
  if (! am2315.begin()) {
     while (1);
  }
  
  while (!Serial) {
    ;
  }
}

void loop() {
  // Wait on Serial Input to turn LED on
  if (Serial.available() > 0){
    int module_number = Serial.parseInt();
    switch (module_number){
      case 1:
        LedOn(module1);
        break;
      case 2:
        LedOn(module2);
        break;
      case 3:
        LedOn(module3);
        break;
      case 4:
        LedOn(module4);
        break;
      case 5:
        LedOn(module5);
        break;
      case 6:
        LedOn(module6);
        break;
      //Numbers for global off and on
      case 80:
        SwitchOn();
        ledOn=true;
        break;
      case 81:
        SwitchOff();
        ledOn=false;
        break;
      // Just a random number to signalise to turn off any led -> 99
      case 90:
        StartPump();
        break;
       case 91:
        StopPump();
        break;
      case 99:
        LedOff();
      default:
        break;
    }
  }


  if(millis() > time_now + printPeriod){
    time_now = millis();
    struct temphum temphum_instance = getHumTemp();
    int lightIntensity = getLightIntensity();
    long distanceToW = getDistance();
    float waterTemp = getWaterTemp();
    float waterpH = getpH();
    Serial.println((String)temphum_instance.temp+","+lightIntensity+","+(String)temphum_instance.humi+","+(String)distanceToW+","+(String)waterTemp+","+(String)waterpH);

  }

  if (blinkLed){
    digitalWrite(ledPin, LOW);
    delay(500);
    digitalWrite(ledPin, HIGH);
    delay(500);
  }

}

//TestingFor phNode sending Stuff



void LedOn(int pin_number){
  ledPin = pin_number;
  blinkLed = true;
}

void LedOff(){
  blinkLed = false;
  if(ledOn){
    digitalWrite(ledPin, HIGH); }
    else{
      digitalWrite(ledPin, LOW);
    }
    }


void SwitchOff(){
    digitalWrite(module1, LOW);
    digitalWrite(module2, LOW);
    digitalWrite(module3, LOW);
    digitalWrite(module4, LOW);
    digitalWrite(module5, LOW);
    digitalWrite(module6, LOW);
 }


 void SwitchOn(){
    digitalWrite(module1, HIGH);
    digitalWrite(module2, HIGH);
    digitalWrite(module3, HIGH);
    digitalWrite(module4, HIGH);
    digitalWrite(module5, HIGH);
    digitalWrite(module6, HIGH);
  }


int getLightIntensity(){
  int photocellReading = analogRead(photocellPin);
  return photocellReading;
}

long getDistance() {
  digitalWrite(trigPin, LOW);
  delayMicroseconds(5);
  digitalWrite(trigPin, HIGH);
  delayMicroseconds(10);
  digitalWrite(trigPin, LOW);

  pinMode(echoPin, INPUT);
  duration = pulseIn(echoPin, HIGH);

  //Returns distance in CM
  return (duration/2) / 29.1;
 }


 void StartPump(){
   sender.sendTriState("000FF0FFFF0F");
   delay(100);
 }
 void StopPump(){
   sender.sendTriState("000FF0FFFFF0");
   delay(100);
 }

float getWaterTemp(){
  sensors.requestTemperatures();
  return sensors.getTempCByIndex(0);
}

 struct temphum getHumTemp(){
  struct temphum temphum_instance;
  am2315.readTemperatureAndHumidity(&temphum_instance.temp, &temphum_instance.humi);
  return temphum_instance;
}
void serialEvent3() {                                 //if the hardware serial port_3 receives a char
  sensorstring = Serial3.readStringUntil(13);         //read the string until we see a <CR>
  sensor_string_complete = true;                      //set the flag used to tell if we have received a completed string from the PC
}
//TODO code is only getting new pH after its read the old(I think), fix so the code gets new Values when it reads

float getpH(){

  if(sensor_string_complete == true){
    pH = sensorstring.toFloat();
  }
  sensorstring="";
  sensor_string_complete = false;
  return pH;
}
