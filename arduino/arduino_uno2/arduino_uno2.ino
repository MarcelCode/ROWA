// Led Pins for Modules
const int module1 = 13;
const int module2 = 11;
const int module3 = 10;
const int module4 = 9;
const int module5 = 8;
const int module6 = 7;

const int tempPin = 0;
int photocellPin = 1;

// Dynamic Variables for LED blink
int ledPin;
bool blinkLed;

// Variables for Sensors
int printPeriod = 1000 * 60 * 5; // every Minute
unsigned long time_now = 0;


void setup() {
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

  Serial.begin(9600);
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
      // Just a random number to signalise to turn off any led -> 99
      case 99:
        LedOff();
      default:
        break;
    }
  }


  if(millis() > time_now + printPeriod){
    time_now = millis();
    float temp = getTemperature();
    int lightIntensity = getLightIntensity();
    Serial.println((String)temp+","+lightIntensity);

  }

  if (blinkLed){
    digitalWrite(ledPin, LOW);
    delay(500);
    digitalWrite(ledPin, HIGH);
    delay(500);
  }

}

void LedOn(int pin_number){
  ledPin = pin_number;
  blinkLed = true;
}

void LedOff(){
  blinkLed = false;
}

float getTemperature(){
  float voltage = analogRead(tempPin) * 0.004882814;
  return (voltage - 0.5) * 100.0;
}

int getLightIntensity(){
  int photocellReading = analogRead(photocellPin);
  return photocellReading;
}
