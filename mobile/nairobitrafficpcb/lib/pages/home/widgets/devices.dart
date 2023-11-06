import 'package:animations/animations.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:flutter_svg/svg.dart';
import 'package:nairobitrafficpcb/model/ble_device.dart';
import 'package:nairobitrafficpcb/pages/control_panel/control_panel_page.dart';

class Devices extends StatelessWidget {
  final BleDevice device;
  final Function(bool) onChanged;

  const Devices({
    Key? key,
    required this.device,
    required this.onChanged,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return OpenContainer(
        transitionType: ContainerTransitionType.fadeThrough,
        transitionDuration: const Duration(milliseconds: 600),
        closedElevation: 0,
        openElevation: 0,
        openShape: const RoundedRectangleBorder(
            borderRadius: BorderRadius.all(Radius.circular(20.0))),
        closedShape: const RoundedRectangleBorder(
            borderRadius: BorderRadius.all(Radius.circular(20.0))),
        openBuilder: (BuildContext context, VoidCallback _) {
          return ControlPanelPage(
            device: device,
          );
        },
        tappable: true,
        closedBuilder: (BuildContext _, VoidCallback openContainer) {
          return AnimatedContainer(
            duration: const Duration(milliseconds: 300),
            decoration: BoxDecoration(
              borderRadius: const BorderRadius.all(
                Radius.circular(20.0),
              ),
              border: Border.all(
                color: Colors.grey[300]!,
                width: 0.6,
              ),
              color: device.isActive ? const Color(0XFF7739FF) : Colors.white,
            ),
            child: Padding(
              padding: const EdgeInsets.all(14.0),
              child: Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                mainAxisAlignment: MainAxisAlignment.spaceBetween,
                children: [
                  Column(
                    crossAxisAlignment: CrossAxisAlignment.start,
                    children: [
                      SvgPicture.asset(
                        'assets/svg/light.svg',
                        height: 30,
                      ),
                      const SizedBox(
                        height: 14,
                      ),
                      SizedBox(
                        width: 65,
                        child: Text(
                          device.device.name,
                          style: TextStyle(
                            height: 1.2,
                            fontSize: 14,
                            color:
                                device.isActive ? Colors.white : Colors.black,
                            fontWeight: FontWeight.w500,
                          ),
                        ),
                      ),
                    ],
                  ),
                  Transform.scale(
                    alignment: Alignment.center,
                    scaleY: 0.8,
                    scaleX: 0.85,
                    child: CupertinoSwitch(
                      onChanged: onChanged,
                      value: device.isActive,
                      activeColor: device.isActive
                          ? Colors.white.withOpacity(0.4)
                          : Colors.black,
                      trackColor: Colors.black,
                    ),
                  ),
                ],
              ),
            ),
          );
        });
  }
}
