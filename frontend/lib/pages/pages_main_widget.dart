import 'package:flutter/cupertino.dart';

import 'package:offline_project/pages/home_page.dart';
import 'package:offline_project/pages/profile_page.dart';

class PagesMainWidget extends StatelessWidget {
  static const String id = "pages_main_widget";

  const PagesMainWidget({super.key});

  @override
  Widget build(BuildContext context) {
    return CupertinoTabScaffold(
      tabBar: CupertinoTabBar(
        items: const <BottomNavigationBarItem>[
          BottomNavigationBarItem(
            icon: Icon(CupertinoIcons.home),
          ),
          BottomNavigationBarItem(
            icon: Icon(CupertinoIcons.person_alt_circle_fill),
          ),
        ],
      ),
      tabBuilder: (BuildContext context, int index) {
        return switch (index) {
          0 => CupertinoTabView(
              builder: (context) => const CupertinoPageScaffold(
                child: HomePage(),
              ),
            ),
          1 => CupertinoTabView(
              builder: (context) => const CupertinoPageScaffold(
                child: ProfilePage(),
              ),
            ),
          _ => throw Exception('Invalid index $index'),
        };
      },
    );
  }
}
