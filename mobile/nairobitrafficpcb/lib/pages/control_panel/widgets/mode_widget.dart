import 'package:flutter/material.dart';
import 'package:nairobitrafficpcb/widgets/transparent_card.dart';

class ModeWidget extends StatelessWidget {
  final int mode;
  final Function(int) changeMode;

  const ModeWidget({Key? key, required this.mode, required this.changeMode})
      : super(key: key);

  @override
  Widget build(BuildContext context) {
    return TransparentCard(
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          const Text(
            "Mode",
            style: TextStyle(
              fontSize: 15,
              color: Colors.white,
              fontWeight: FontWeight.w500,
            ),
          ),
          const SizedBox(
            height: 5,
          ),
          Row(
            mainAxisAlignment: MainAxisAlignment.spaceBetween,
            children: [
              _button(1, mode == 1),
              _button(2, mode == 2),
              _button(3, mode == 3),
              _button(4, mode == 4),
              _button(5, mode == 5),
              _button(6, mode == 6),
              _button(7, mode == 7),
            ],
          ),
        ],
      ),
    );
  }

  ElevatedButton _button(int mode, bool isActive) {
    return ElevatedButton(
      style: ElevatedButton.styleFrom(
          onPrimary: isActive ? Colors.black : Colors.white,
          primary: isActive ? Colors.white : Colors.transparent,
          minimumSize: const Size(38, 38),
          padding: EdgeInsets.zero,
          shape: const CircleBorder(),
          side: BorderSide(
            color: Colors.white.withOpacity(0.4),
          ),
          elevation: 0),
      onPressed: () => changeMode(mode),
      child: Text(mode.toString()),
    );
  }
}
