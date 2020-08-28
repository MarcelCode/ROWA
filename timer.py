import time
import serial

ser = serial.Serial('/dev/ttyACM0', 9600)  # open serial port
print(ser.name)         # check which port was really used
ser.write(b'hello')     # write a string
ser.close()             # close port

ON = 18 * 60 + 48 # turn on at half 8 in the evening
OFF = 6 * 60 + 15 # turn off at 6:15 in the morning

switched_on = False # hold current state

while True:
    print("Starting timer..")
    stm = time.localtime()
    now = stm.tm_hour * 60 + stm.tm_min
    if (now - ON) % 1440 < (OFF - ON) % 1440:
        should_be_on = True
    else:
        should_be_on = False
    if should_be_on and not switched_on:
        ser.write('70')
        time.sleep(60.0)
        ser.write('90')
        print("turned on")

    elif switched_on and not should_be_on:
        ser.write('91')
        time.sleep(5.0)
        ser.write('71')
        print("turned off")
        switched_on = False
    time.sleep(10.0)