import 'package:flutter/material.dart';

class CustomBottomNavigationBar extends StatefulWidget {
  final PageController pageController_;

  const CustomBottomNavigationBar({
    super.key,
    required this.pageController_,
  });

  @override
  State<CustomBottomNavigationBar> createState() =>
      _CustomBottomNavigationBarState();
}

class _CustomBottomNavigationBarState extends State<CustomBottomNavigationBar> {
  int selectedIndex_ = 0;

  final List<BottomNavigationBarItem> icons_ = const [
    BottomNavigationBarItem(
      backgroundColor: Colors.black,
      icon: Icon(
        Icons.home_filled,
        size: 35,
      ),
      label: '',
    ),
    BottomNavigationBarItem(
      icon: Icon(
        Icons.person_outline,
        size: 35,
      ),
      label: '',
    ),
  ];

  void processTap(int index) {
    setState(() {
      selectedIndex_ = index;
      widget.pageController_.animateToPage(
        index,
        duration: const Duration(milliseconds: 100),
        curve: Curves.ease,
      );
    });
  }

  @override
  Widget build(BuildContext context) {
    return BottomNavigationBar(
      onTap: processTap,
      showSelectedLabels: false,
      showUnselectedLabels: false,
      currentIndex: selectedIndex_,
      items: icons_,
    );
  }
}
