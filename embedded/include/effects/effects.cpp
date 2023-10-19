#include "effects.h"

Effect::Effect(Pins pins)
{
    this->pins = pins;
}
Effect::~Effect()
{
    this->pins = {};
}
void Effect::SetColor(uint32_t color, Road road, uint8_t brightness)
{
    
}
void Effect::SetEffect(EffectTypes effect, Road road, uint8_t brightness)
{
}
