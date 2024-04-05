import 'package:flutter/material.dart';
import 'package:frontend/components/button.dart';
import 'package:frontend/components/text_field.dart';

class RegisterPage extends StatefulWidget {
//  final Function()? onTap;

  const RegisterPage({super.key});

  @override
  State<RegisterPage> createState() => _RegisterPageState();
}

class _RegisterPageState extends State<RegisterPage>
    with SingleTickerProviderStateMixin {
  late AnimationController _controller;

  final emailTextController = TextEditingController();
  final passwordTextController = TextEditingController();
  final confirmPasswordTextController = TextEditingController();

  @override
  void initState() {
    super.initState();
    _controller = AnimationController(vsync: this);
  }

  @override
  void dispose() {
    _controller.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: Colors.grey[300],
      body: SafeArea(
        child: Center(
            child: Padding(
          padding: const EdgeInsets.symmetric(horizontal: 25.0),
          child: Column(
            mainAxisAlignment: MainAxisAlignment.center,
            children: [
              const Icon(
                Icons.lock,
                size: 50,
              ),
              const SizedBox(
                height: 50,
              ),
              Text(
                "Registration",
                style: TextStyle(
                  color: Colors.grey[700],
                ),
              ),
              const SizedBox(
                height: 25,
              ),
              MyTextField(
                  controller: emailTextController,
                  hintText: 'Email',
                  obscureText: false),
              const SizedBox(
                height: 10,
              ),
              MyTextField(
                  controller: passwordTextController,
                  hintText: 'Password',
                  obscureText: true),
              const SizedBox(
                height: 10,
              ),
              MyTextField(
                  controller: confirmPasswordTextController,
                  hintText: 'Confirm password',
                  obscureText: true),
              const SizedBox(
                height: 10,
              ),
              MyButton(
                onTap: () {},
                text: 'Sign Up',
              ),
              const SizedBox(
                height: 25,
              ),
              Row(
                mainAxisAlignment: MainAxisAlignment.center,
                children: [
                  Text(
                    "Already have an account?",
                    style: TextStyle(
                      color: Colors.grey[700],
                    ),
                  ),
                  const SizedBox(
                    width: 10,
                  ),
                  GestureDetector(
                    onTap: () {}, //widget.onTap,
                    child: const Text(
                      "Login now",
                      style: TextStyle(
                        fontWeight: FontWeight.bold,
                        color: Colors.blue,
                      ),
                    ),
                  ),
                ],
              ),
            ],
          ),
        )),
      ),
    );
  }
}
