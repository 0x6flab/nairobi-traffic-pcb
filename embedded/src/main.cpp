#include "leds.h"
#include "Arduino.h"

Led Kiambu;

void setup()
{
    Serial.begin(115200);
    Kiambu = Led(15, 6, 58);
}

void loop()
{
    Serial.println("Hello World");
    delay(1000);
    Kiambu.run();
}
