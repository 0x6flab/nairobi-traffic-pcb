import 'package:flutter/material.dart';
import 'package:google_fonts/google_fonts.dart';
import 'package:nairobitrafficpcb/pages/home/home_page.dart';

void main() {
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Nairobi Traffic PCB',
      debugShowCheckedModeBanner: false,
      theme: ThemeData(
        fontFamily: GoogleFonts.poppins().fontFamily,
        sliderTheme: const SliderThemeData(
          trackShape: RectangularSliderTrackShape(),
          trackHeight: 2.5,
          thumbShape: RoundSliderThumbShape(enabledThumbRadius: 8.0),
          overlayShape: RoundSliderOverlayShape(overlayRadius: 15.0),
        ),
      ),
      home: const HomePage(),
    );
  }
}
