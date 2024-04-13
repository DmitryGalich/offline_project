import 'package:flutter/material.dart';

import 'package:frontend/components/bottom_navigation_bar.dart';
import 'package:frontend/pages/home/home_page.dart';
import 'package:frontend/pages/profile/profile_page.dart';

import 'package:frontend/services/web_socket_client.dart';

final webSocketClient = WebSocketClient();

class RootPage extends StatelessWidget {
  const RootPage({super.key});

  @override
  Widget build(BuildContext context) {
    final PageController pageController_ = PageController(initialPage: 0);

    const List<Widget> pages_ = [
      HomePage(),
      ProfilePage(),
    ];

    return Scaffold(
      body: PageView(
        controller: pageController_,
        physics: const NeverScrollableScrollPhysics(),
        children: pages_,
      ),
      bottomNavigationBar: CustomBottomNavigationBar(
        pageController_: pageController_,
      ),
    );
  }
}
