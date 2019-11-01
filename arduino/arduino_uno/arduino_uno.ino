// Led Pins for Modules
const int module1 = 12;
const int module2 = 11;
const int module3 = 10;
const int module4 = 9;
const int module5 = 8;
const int module6 = 7;

const int tempPin = 0;

// Dynamic Variables for LED blink
int ledPin;
int ledState = LOW;

// Variables for Temp
int printPeriod = 2000;
unsigned long time_now = 0;


void setup() {
  pinMode(module1, OUTPUT);
  pinMode(module2, OUTPUT);
  pinMode(module3, OUTPUT);
  pinMode(module4, OUTPUT);
  pinMode(module5, OUTPUT);
  pinMode(module6, OUTPUT);
  
  Serial.begin(9600);
  while (!Serial) {
    ;
  }

  // Print Temperature
  
}

void loop() {
  // Wait on Serial Input to turn LED on
  if (Serial.available() > 0){
    int pin_number = Serial.parseInt();
    switch (pin_number){
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
      case 99:
        LedOff();
      default:
        break;
    }
  }

  
  if(millis() > time_now + printPeriod){
    time_now = millis();
    float temp = getTemperature();
    Serial.println(temp);
  }
}

void LedOn(int pin_number){
  if (ledState){
        ledState = LOW;
        digitalWrite(ledPin, ledState);
      }
      ledPin = pin_number;
      ledState = HIGH;
      digitalWrite(ledPin, ledState);
}

void LedOff(){
  ledState = LOW;
  digitalWrite(ledPin, ledState);
}

float getTemperature(){
  float voltage = analogRead(tempPin) * 0.004882814;
  return (voltage - 0.5) * 100.0;
}
