import 'package:flutter_blue/flutter_blue.dart';

const String serviceUUID = "4fafc201-1fb5-459e-8fcc-c5c9c331914b";
const String characteristicUUID = "beb5483e-36e1-4688-b7f5-ea07361b26a8";

class BleDevice {
  bool isActive = false;
  BluetoothDevice device;

  BleDevice({required this.isActive, required this.device});

  Future<void> write(String data) async {
    List<BluetoothService> services = await device.discoverServices();
    for (var service in services) {
      if (service.uuid.toString() == serviceUUID) {
        for (var characteristic in service.characteristics) {
          if (characteristic.uuid.toString() == characteristicUUID) {
            final hexData = data.split('').map((e) => e.codeUnitAt(0)).toList();
            characteristic.write(hexData, withoutResponse: true);
          }
        }
      }
    }
  }
}
