#ifndef LIB_LEDS_LEDS_H_
#define LIB_LEDS_LEDS_H_

#include <WS2812FX.h>

/*
Led is a class that represents a single LED strip.

It is a wrapper around the WS2812FX library.
*/
class Led

{
private:
    WS2812FX _pin;
    uint8_t brightness = 100;
    uint16_t speed = 255;

public:
    /*
    Led is a class that represents a single LED strip.
    It is a wrapper around the WS2812FX library.

     * @param count: Number of LEDs in the strip
     * @param gpio: GPIO pin number
     * @param brightness: Brightness value between 0 and 255
     * @param speed: Speed value between 0 and 1000
     */
    Led(uint16_t count, uint8_t gpio, uint8_t mode);
    void run();
};

#endif /* LIB_LEDS_LEDS_H_ */