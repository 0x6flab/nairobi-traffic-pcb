import 'package:flutter/material.dart';
import 'package:nairobitrafficpcb/widgets/transparent_card.dart';

class BrightnessWidget extends StatelessWidget {
  final double brightness;
  final Function(double) changeBrightness;
  final Function(double) writeBrightness;

  const BrightnessWidget({
    Key? key,
    required this.brightness,
    required this.changeBrightness,
    required this.writeBrightness,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Center(
      child: TransparentCard(
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          mainAxisAlignment: MainAxisAlignment.spaceBetween,
          children: [
            const Text(
              "Brightness",
              style: TextStyle(
                  fontSize: 15,
                  color: Colors.white,
                  fontWeight: FontWeight.w500),
            ),
            const SizedBox(
              height: 5,
            ),
            Row(
              mainAxisAlignment: MainAxisAlignment.spaceBetween,
              children: [
                const Text(
                  '0%',
                  style: TextStyle(color: Colors.white),
                ),
                Expanded(
                  child: Slider(
                    min: 0,
                    max: 100,
                    value: brightness,
                    activeColor: Colors.white,
                    inactiveColor: Colors.white30,
                    onChanged: changeBrightness,
                    onChangeEnd: writeBrightness,
                  ),
                ),
                const Text(
                  '100%',
                  style: TextStyle(color: Colors.white),
                ),
              ],
            ),
            const SizedBox(
              height: 12,
            ),
          ],
        ),
      ),
    );
  }
}
