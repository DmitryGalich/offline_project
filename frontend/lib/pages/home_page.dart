import 'package:flutter/cupertino.dart';

class HomePage extends StatelessWidget {
  static const String id = "home_page";

  const HomePage({super.key});

  @override
  Widget build(BuildContext context) {
    return const CupertinoPageScaffold(
      navigationBar: CupertinoNavigationBar(
        middle: Text('Home'),
      ),
      child: SizedBox(),
    );
  }
}
