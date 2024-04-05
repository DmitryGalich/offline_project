import 'package:flutter/material.dart';

class BottomNavBarCustom extends StatelessWidget {
  const BottomNavBarCustom({
    super.key,
    required this.pageController,
  });
  final PageController pageController;
  @override
  Widget build(BuildContext context) {
    return BottomNavigationBar(
      onTap: (index) {
        pageController.jumpToPage(index);
      },
      showSelectedLabels: false,
      showUnselectedLabels: false,
      items: const [
        BottomNavigationBarItem(
          backgroundColor: Colors.black,
          icon: Icon(
            Icons.home_filled,
            size: 35,
          ),
          label: '',
        ),
        BottomNavigationBarItem(
          backgroundColor: Colors.black,
          icon: Icon(
            Icons.person_outline,
            size: 35,
          ),
          label: '',
        ),
      ],
    );
  }
}
