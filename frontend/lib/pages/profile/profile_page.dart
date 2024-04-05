import 'package:flutter/material.dart';

class ProfilePage extends StatelessWidget {
  const ProfilePage({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Profile page'),
      ),
      body: const SafeArea(
        child: Center(
          child: Text(
            'Profile page content',
            style: TextStyle(fontSize: 20),
          ),
        ),
      ),
    );
  }
}
