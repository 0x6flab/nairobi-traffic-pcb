import 'package:flutter/material.dart';
import 'package:flutter_circle_color_picker/flutter_circle_color_picker.dart';
import 'package:nairobitrafficpcb/model/ble_device.dart';
import 'package:nairobitrafficpcb/pages/control_panel/widgets/brightness_widget.dart';
import 'package:nairobitrafficpcb/pages/control_panel/widgets/mode_widget.dart';
import 'package:nairobitrafficpcb/widgets/custom_appbar.dart';

class ControlPanelPage extends StatefulWidget {
  final BleDevice device;

  const ControlPanelPage({Key? key, required this.device}) : super(key: key);
  @override
  State<ControlPanelPage> createState() => _ControlPanelPageState();
}

class _ControlPanelPageState extends State<ControlPanelPage>
    with TickerProviderStateMixin {
  bool isActive = false;
  int mode = 1;
  double brightness = 50;
  Color currentColor = const Color(0XFF7739FF).withOpacity(0.5);

  final _controller = CircleColorPickerController(
    initialColor: const Color(0XFF7739FF).withOpacity(0.5),
  );

  @override
  void initState() {
    super.initState();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Container(
        width: MediaQuery.of(context).size.width,
        height: MediaQuery.of(context).size.height,
        decoration: BoxDecoration(
          gradient: LinearGradient(
              begin: Alignment.topCenter,
              end: Alignment.bottomCenter,
              colors: <Color>[
                Colors.white,
                currentColor.withOpacity(brightness / 100),
                currentColor,
              ]),
        ),
        child: Padding(
          padding: const EdgeInsets.fromLTRB(15, 50, 15, 0),
          child: Padding(
            padding: const EdgeInsets.all(8.0),
            child: Column(
              children: [
                const CustomAppBar(),
                const SizedBox(
                  height: 20,
                ),
                Expanded(
                  child: Column(
                    mainAxisAlignment: MainAxisAlignment.spaceBetween,
                    children: [
                      const SizedBox(
                        height: 20,
                      ),
                      Center(
                        child: CircleColorPicker(
                          controller: _controller,
                          size: const Size(300, 300),
                          strokeWidth: 8,
                          thumbSize: 48,
                          onChanged: (color) {
                            setState(() => currentColor = color);
                          },
                          onEnded: (value) {
                            final hexValue = value.value.toRadixString(16);
                            final data = 'c 0x$hexValue';
                            widget.device.write(data);
                          },
                        ),
                      ),
                      controls(),
                    ],
                  ),
                )
              ],
            ),
          ),
        ),
      ),
    );
  }

  Widget controls() {
    return Column(
      children: [
        Row(
          mainAxisAlignment: MainAxisAlignment.spaceBetween,
          children: [
            Expanded(
              child: BrightnessWidget(
                  brightness: brightness,
                  writeBrightness: (val) {
                    final data = 'b ${val.toInt()}';
                    widget.device.write(data);
                  },
                  changeBrightness: (val) => setState(() {
                        brightness = val;
                      })),
            ),
          ],
        ),
        const SizedBox(
          height: 15,
        ),
        ModeWidget(
          mode: mode,
          changeMode: (val) => setState(() {
            mode = val;
            final data = 'm ${val.toInt()}';
            widget.device.write(data);
          }),
        ),
        const SizedBox(
          height: 15,
        ),
      ],
    );
  }
}
