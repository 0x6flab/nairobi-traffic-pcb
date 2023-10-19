#include "leds.h"

Led::Led(uint16_t count, uint8_t gpio, uint8_t mode) : _pin(count, gpio, NEO_RGB + NEO_KHZ800)
{
    this->_pin.init();
    this->_pin.setBrightness(this->brightness);
    this->_pin.setSpeed(this->speed);
    this->_pin.setMode(mode);
    this->_pin.start();
}

void Led::run()
{

    this->_pin.service();
}