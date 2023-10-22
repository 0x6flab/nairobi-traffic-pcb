#include <BLEDevice.h>
#include <BLEUtils.h>
#include <BLEServer.h>
#include <Arduino.h>
#include <WS2812FX.h>
#include <Preferences.h>

#define SERVICE_UUID "4fafc201-1fb5-459e-8fcc-c5c9c331914b"
#define CHARACTERISTIC_UUID "beb5483e-36e1-4688-b7f5-ea07361b26a8"
#define PROJECT_NAME "Nairobi Traffic PCB"
#define MESSAGE "Hello this is a message from the PCB"

#define THIKA_COUNT 15
#define THIKA_LED_PIN 6
#define KIAMBU_COUNT 3
#define KIAMBU_LED_PIN 4
#define JUJA_COUNT 9
#define JUJA_LED_PIN 1
#define JOGOO_COUNT 5
#define JOGOO_LED_PIN 2
#define KANGUNDO_COUNT 3
#define KANGUNDO_LED_PIN 39
#define KODHEK_COUNT 4
#define KODHEK_LED_PIN 12
#define DAGORETTI_COUNT 3
#define DAGORETTI_LED_PIN 21
#define NGONG_COUNT 6
#define NGONG_LED_PIN 7
#define LANGATA_COUNT 7
#define LANGATA_LED_PIN 5
#define MOMBASA_COUNT 7
#define MOMBASA_LED_PIN 8
#define EMBAKASI_COUNT 4
#define EMBAKASI_LED_PIN 13
#define KAYOLE_COUNT 5
#define KAYOLE_LED_PIN 11
#define KILELESHWA_COUNT 4
#define KILELESHWA_LED_PIN 38
#define WAYAKI_COUNT 4
#define WAYAKI_LED_PIN 10
#define PARKLANDS_COUNT 3
#define PARKLANDS_LED_PIN 9
#define LIMURU_COUNT 6
#define LIMURU_LED_PIN 14

WS2812FX LEDS[16] = {
    WS2812FX(THIKA_COUNT, THIKA_LED_PIN, NEO_RGB + NEO_KHZ800),
    WS2812FX(KIAMBU_COUNT, KIAMBU_LED_PIN, NEO_RGB + NEO_KHZ800),
    WS2812FX(JUJA_COUNT, JUJA_LED_PIN, NEO_RGB + NEO_KHZ800),
    WS2812FX(JOGOO_COUNT, JOGOO_LED_PIN, NEO_RGB + NEO_KHZ800),
    WS2812FX(KANGUNDO_COUNT, KANGUNDO_LED_PIN, NEO_RGB + NEO_KHZ800),
    WS2812FX(KODHEK_COUNT, KODHEK_LED_PIN, NEO_RGB + NEO_KHZ800),
    WS2812FX(DAGORETTI_COUNT, DAGORETTI_LED_PIN, NEO_RGB + NEO_KHZ800),
    WS2812FX(NGONG_COUNT, NGONG_LED_PIN, NEO_RGB + NEO_KHZ800),
    WS2812FX(LANGATA_COUNT, LANGATA_LED_PIN, NEO_RGB + NEO_KHZ800),
    WS2812FX(MOMBASA_COUNT, MOMBASA_LED_PIN, NEO_RGB + NEO_KHZ800),
    WS2812FX(EMBAKASI_COUNT, EMBAKASI_LED_PIN, NEO_RGB + NEO_KHZ800),
    WS2812FX(KAYOLE_COUNT, KAYOLE_LED_PIN, NEO_RGB + NEO_KHZ800),
    WS2812FX(KILELESHWA_COUNT, KILELESHWA_LED_PIN, NEO_RGB + NEO_KHZ800),
    WS2812FX(WAYAKI_COUNT, WAYAKI_LED_PIN, NEO_RGB + NEO_KHZ800),
    WS2812FX(PARKLANDS_COUNT, PARKLANDS_LED_PIN, NEO_RGB + NEO_KHZ800),
    WS2812FX(LIMURU_COUNT, LIMURU_LED_PIN, NEO_RGB + NEO_KHZ800),
};

Preferences preferences;

void process_command(char *scmd);

class Callback : public BLECharacteristicCallbacks
{
  void onWrite(BLECharacteristic *pCharacteristic)
  {
    std::string value = pCharacteristic->getValue();
    process_command((char *)value.c_str());
    Serial.println(value.c_str());
  }
};

void setup()
{
  Serial.begin(115200);
  preferences.begin("ntpcb-ledmode", false);
  uint8_t mode = preferences.getUChar("mode", 0);
  uint32_t color = preferences.getUInt("color", 0x007BFF);
  uint8_t brightness = preferences.getUChar("brightness", 30);
  uint16_t speed = preferences.getUShort("speed", 1000);

  for (int i = 0; i < 16; i++)
  {
    LEDS[i].init();
    LEDS[i].setBrightness(brightness);
    LEDS[i].setSpeed(speed);
    LEDS[i].setColor(color);
    LEDS[i].setMode(mode);
    LEDS[i].start();
  }

  BLEDevice::init(PROJECT_NAME);
  BLEServer *pServer = BLEDevice::createServer();
  BLEService *pService = pServer->createService(SERVICE_UUID);
  BLECharacteristic *pCharacteristic = pService->createCharacteristic(
      CHARACTERISTIC_UUID,
      BLECharacteristic::PROPERTY_WRITE);

  pCharacteristic->setCallbacks(new Callback());
  pCharacteristic->setValue(MESSAGE);
  pService->start();

  BLEAdvertising *pAdvertising = BLEDevice::getAdvertising();
  pAdvertising->addServiceUUID(SERVICE_UUID);
  pAdvertising->setScanResponse(true);
  pAdvertising->setMinPreferred(0x06); // functions that help with iPhone connections issue
  pAdvertising->setMinPreferred(0x12);

  BLEDevice::startAdvertising();
}

void loop()
{
  for (int i = 0; i < 16; i++)
  {
    LEDS[i].service();
  }
}

void process_command(char *scmd)
{
  if (strcmp(scmd, "b+") == 0)
  {
    for (int i = 0; i < 16; i++)
    {
      LEDS[i].increaseBrightness(25);
    }
  }

  if (strcmp(scmd, "b-") == 0)
  {
    for (int i = 0; i < 16; i++)
    {
      LEDS[i].decreaseBrightness(25);
    }
  }

  if (strncmp(scmd, "b ", 2) == 0)
  {
    uint8_t b = (uint8_t)atoi(scmd + 2);
    for (int i = 0; i < 16; i++)
    {
      LEDS[i].setBrightness(b);
    }
  }

  if (strcmp(scmd, "s+") == 0)
  {
    for (int i = 0; i < 16; i++)
    {
      LEDS[i].setSpeed(LEDS[i].getSpeed() * 1.2);
    }
  }

  if (strcmp(scmd, "s-") == 0)
  {
    for (int i = 0; i < 16; i++)
    {
      LEDS[i].setSpeed(LEDS[i].getSpeed() * 0.8);
    }
  }

  if (strncmp(scmd, "s ", 2) == 0)
  {
    uint16_t s = (uint16_t)atoi(scmd + 2);
    for (int i = 0; i < 16; i++)
    {
      LEDS[i].setSpeed(s);
    }
  }

  if (strncmp(scmd, "m ", 2) == 0)
  {
    uint8_t m = (uint8_t)atoi(scmd + 2);
    for (int i = 0; i < 16; i++)
    {
      LEDS[i].setMode(m);
    }
  }

  if (strncmp(scmd, "c ", 2) == 0)
  {
    uint32_t c = (uint32_t)strtoul(scmd + 2, NULL, 16);
    for (int i = 0; i < 16; i++)
    {
      LEDS[i].setColor(c);
    }
  }

  preferences.putUChar("mode", LEDS[0].getMode());
  preferences.putUInt("color", LEDS[0].getColor());
  preferences.putUChar("brightness", LEDS[0].getBrightness());
  preferences.putUShort("speed", LEDS[0].getSpeed());
}
