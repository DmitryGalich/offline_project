import 'package:flutter/cupertino.dart';

class LoginPage extends StatelessWidget {
  static const String id = "login_page";

  @override
  Widget build(BuildContext context) {
    return const CupertinoPageScaffold(
      navigationBar: CupertinoNavigationBar(
        middle: Text('Cupertino Store'),
      ),
      child: SizedBox(),
    );
  }
}
