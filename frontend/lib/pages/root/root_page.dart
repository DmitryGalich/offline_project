import 'package:flutter/material.dart';

import 'package:frontend/components/bottom_navigation_bar.dart';
import 'package:frontend/pages/home/home_page.dart';
import 'package:frontend/pages/profile/profile_page.dart';

class RootPage extends StatelessWidget {
  const RootPage({super.key});

  @override
  Widget build(BuildContext context) {
    final PageController pageController_ = PageController(initialPage: 0);

    return Scaffold(
      body: PageView(
        controller: pageController_,
        physics: const NeverScrollableScrollPhysics(),
        children: const [
          HomePage(),
          ProfilePage(),
        ],
      ),
      bottomNavigationBar: CustomBottomNavigationBar(
        pageController_: pageController_,
      ),
    );
  }
}
