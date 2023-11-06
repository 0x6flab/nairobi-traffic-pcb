import 'package:flutter/material.dart';
import 'package:flutter_blue/flutter_blue.dart';
import 'package:get/get.dart';
import 'package:nairobitrafficpcb/controller/ble_controller.dart';
import 'package:nairobitrafficpcb/pages/home/widgets/devices.dart';
import 'package:nairobitrafficpcb/model/ble_device.dart';

class HomePage extends StatefulWidget {
  const HomePage({Key? key}) : super(key: key);

  @override
  State<HomePage> createState() => _HomePageState();
}

class _HomePageState extends State<HomePage> {
  List<BleDevice> devices = [];

  bool addDevice(ScanResult scanResult) {
    if ((scanResult.device.name != "") &&
        (scanResult.device.id.toString() != "")) {
      devices.add(
        BleDevice(
          device: scanResult.device,
          isActive: false,
        ),
      );
      return true;
    }
    return false;
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: GetBuilder<BleController>(
        init: BleController(),
        builder: (controller) {
          return Container(
            width: MediaQuery.of(context).size.width,
            height: MediaQuery.of(context).size.height,
            decoration: const BoxDecoration(
              gradient: LinearGradient(
                begin: Alignment.topLeft,
                end: Alignment.bottomRight,
                colors: <Color>[
                  Color(0xFFfce2e1),
                  Colors.white,
                ],
              ),
            ),
            child: Padding(
              padding: const EdgeInsets.fromLTRB(20, 15, 20, 20),
              child: SafeArea(
                child: Column(
                  children: [
                    const Row(
                      mainAxisAlignment: MainAxisAlignment.spaceBetween,
                      crossAxisAlignment: CrossAxisAlignment.center,
                      children: [
                        Text(
                          "Hello Nairobian!",
                          style: TextStyle(
                            fontSize: 28,
                            color: Colors.black,
                            fontWeight: FontWeight.bold,
                          ),
                        ),
                      ],
                    ),
                    const SizedBox(
                      height: 30,
                    ),
                    Expanded(
                      child: Container(
                        width: MediaQuery.of(context).size.width,
                        decoration: const BoxDecoration(
                          borderRadius: BorderRadius.only(
                            topRight: Radius.circular(30.0),
                            topLeft: Radius.circular(30.0),
                          ),
                          color: Colors.white,
                        ),
                        child: Padding(
                          padding: const EdgeInsets.all(20.0),
                          child: Column(
                            crossAxisAlignment: CrossAxisAlignment.start,
                            children: [
                              const SizedBox(
                                height: 5,
                              ),
                              Row(
                                mainAxisAlignment:
                                    MainAxisAlignment.spaceBetween,
                                crossAxisAlignment: CrossAxisAlignment.start,
                                children: [
                                  const Column(
                                    crossAxisAlignment:
                                        CrossAxisAlignment.start,
                                    children: [
                                      Text(
                                        "Devices",
                                        style: TextStyle(
                                          fontSize: 15,
                                          color: Colors.grey,
                                          fontWeight: FontWeight.normal,
                                        ),
                                      ),
                                      Text(
                                        "Available Devices",
                                        style: TextStyle(
                                          height: 1.1,
                                          fontSize: 17,
                                          color: Colors.black,
                                          fontWeight: FontWeight.w600,
                                        ),
                                      ),
                                    ],
                                  ),
                                  Icon(
                                    Icons.more_horiz,
                                    color: Colors.grey[300],
                                    size: 30,
                                  )
                                ],
                              ),
                              const SizedBox(
                                height: 10,
                              ),
                              StreamBuilder<List<ScanResult>>(
                                stream: controller.scanResults,
                                builder: (context, snapshot) {
                                  if (snapshot.hasData) {
                                    return Expanded(
                                      child: GridView.builder(
                                          padding: const EdgeInsets.only(
                                            top: 10,
                                            bottom: 20,
                                          ),
                                          gridDelegate:
                                              const SliverGridDelegateWithMaxCrossAxisExtent(
                                            maxCrossAxisExtent: 200,
                                            childAspectRatio: 3 / 4,
                                            crossAxisSpacing: 20,
                                            mainAxisSpacing: 20,
                                          ),
                                          itemCount: snapshot.data!.length,
                                          shrinkWrap: true,
                                          itemBuilder: (context, index) {
                                            if (addDevice(
                                                snapshot.data![index])) {
                                              return Devices(
                                                device: devices[index],
                                                onChanged: (val) {
                                                  handlerConnection(index);
                                                },
                                              );
                                            }
                                            return null;
                                          }),
                                    );
                                  } else {
                                    return ElevatedButton(
                                      onPressed: () => controller.scanDevices(),
                                      child: const Text("Scan"),
                                    );
                                  }
                                },
                              ),
                              ElevatedButton(
                                onPressed: () => controller.scanDevices(),
                                child: const Text("Scan"),
                              ),
                            ],
                          ),
                        ),
                      ),
                    ),
                  ],
                ),
              ),
            ),
          );
        },
      ),
    );
  }

  void handlerConnection(int index) {
    return setState(() {
      devices[index].isActive = !devices[index].isActive;
      if (devices[index].isActive) {
        devices[index].device.connect();
      } else {
        devices[index].device.disconnect();
      }
    });
  }
}
