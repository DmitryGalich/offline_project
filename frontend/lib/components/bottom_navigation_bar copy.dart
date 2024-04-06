import 'package:flutter/material.dart';

class CustomBottomNavigationBar extends StatelessWidget {
  final PageController pageController_;

  const CustomBottomNavigationBar({
    super.key,
    required this.pageController_,
  });

  @override
  Widget build(BuildContext context) {
    return BottomNavigationBar(
      onTap: (index) {
        pageController_.jumpToPage(index);
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
