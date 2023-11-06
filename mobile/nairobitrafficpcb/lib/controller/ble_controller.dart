import 'package:get/get.dart';
import 'package:permission_handler/permission_handler.dart';
import 'package:flutter_blue/flutter_blue.dart';

const scanDuration = Duration(seconds: 10);

class BleController extends GetxController {
  FlutterBlue flutterBlue = FlutterBlue.instance;

  Future scanDevices() async {
    var blePermission = await Permission.bluetoothScan.status;
    if (blePermission.isDenied) {
      if (await Permission.bluetoothScan.request().isGranted) {
        if (await Permission.bluetoothConnect.request().isGranted) {
          flutterBlue.startScan(timeout: scanDuration);
          flutterBlue.stopScan();
        }
      }
    } else {
      flutterBlue.startScan(timeout: scanDuration);
      flutterBlue.stopScan();
    }
  }

  Stream<List<ScanResult>> get scanResults => flutterBlue.scanResults;
}
