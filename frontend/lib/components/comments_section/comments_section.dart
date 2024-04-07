import 'package:flutter/material.dart';

import 'input_section.dart';

class CommentsSection extends StatefulWidget {
  const CommentsSection({super.key});

  @override
  State<CommentsSection> createState() => _CommentsSectionState();
}

class _CommentsSectionState extends State<CommentsSection> {
  String text_ = '1231';

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Center(
        child: Column(
          children: [
            Text(
              text_,
            ),
            const InputSection(),
          ],
        ),
      ),
    );
  }
}
