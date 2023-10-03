import 'package:flutter/cupertino.dart';

import 'package:offline_project/pages/pages_main_widget.dart';
import 'package:offline_project/pages/login_page.dart';

class App extends StatefulWidget {
  const App({super.key});

  @override
  State<App> createState() => _AppState();
}

class _AppState extends State<App> {
  @override
  Widget build(BuildContext context) {
    return CupertinoApp(home: const LoginPage(), routes: {
      LoginPage.id: (context) => const LoginPage(
            key: null,
          ),
      PagesMainWidget.id: (context) => const PagesMainWidget(
            key: null,
          ),
    });
  }
}
