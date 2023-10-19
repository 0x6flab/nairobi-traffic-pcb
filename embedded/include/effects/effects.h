#ifndef INCLUDE_EFFECTS_EFFECTS_H_
#define INCLUDE_EFFECTS_EFFECTS_H_
#include <Arduino.h>

struct Pin
{
    int pin;
    int count;
};

struct Pins
{
    Pin thika;
    Pin kiambu;
    Pin juja;
    Pin jogoo;
    Pin kangundo;
    Pin kodhek;
    Pin dagoretti;
    Pin ngong;
    Pin langata;
    Pin mombasa;
    Pin embakasi;
    Pin kayole;
    Pin kileleshwa;
    Pin wayaki;
    Pin parklands;
    Pin limuru;
};

enum Road
{
    Thika,
    Kiambu,
    Juja,
    Jogoo,
    Kangundo,
    Kodhek,
    Dagoretti,
    Ngong,
    Langata,
    Mombasa,
    Embakasi,
    Kayole,
    Kileleshwa,
    Wayaki,
    Parklands,
    Limuru,
    All
};

enum EffectTypes
{
    Static,
    Blink,
    Breath,
    ColorWipe,
    ColorWipeInv,
    ColorWipeRev,
    ColorWipeRevInv,
    ColorWipeRandom,
    RandomColor,
    SingleDynamic,
    MultiDynamic,
    Rainbow,
    RainbowCycle,
    Scan,
    DualScan,
    Fade,
    TheaterChase,
    TheaterChaseRainbow,
    RunningLights,
    Twinkle,
    TwinkleRandom,
    TwinkleFade,
    TwinkleFadeRandom,
    Sparkle,
    FlashSparkle,
    HyperSparkle,
    Strobe,
    StrobeRainbow,
    MultiStrobe,
    BlinkRainbow,
    ChaseWhite,
    ChaseColor,
    ChaseRandom,
    ChaseRainbow,
    ChaseFlash,
    ChaseFlashRandom,
    ChaseRainbowWhite,
    ChaseBlackout,
    ChaseBlackoutRainbow,
    ColorSweepRandom,
    RunningColor,
    RunningRedBlue,
    RunningRandom,
    LarsonScanner,
    Comet,
    Fireworks,
    FireworksRandom,
    MerryChristmas,
    FireFlicker,
    FireFlickerSoft,
    FireFlickerIntense,
    CircusCombustus,
    Halloween,
    BiColorChase,
    TriColorChase,
    TwinkleFox,
    Rain,
    BlockDissolve,
    Icu,
    DualLarson,
    RunningRandom2,
    FillerUp,
    RainbowLarson,
    RainbowFireworks,
    TriFade,
    VuMeter,
    Heartbeat,
    Bits,
    MultiComet,
    Flipbook,
    Popcorn,
    Oscillator
};

class Effect
{
private:
    Pins pins;

public:
    Effect(Pins pins);
    void SetColor(uint32_t color, Road road, uint8_t brightness);
    void SetEffect(EffectTypes effect, Road road, uint8_t brightness);
    ~Effect();
};

#endif /* INCLUDE_EFFECTS_EFFECTS_H_ */