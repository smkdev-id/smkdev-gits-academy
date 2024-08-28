import '/flutter_flow/flutter_flow_animations.dart';
import '/flutter_flow/flutter_flow_theme.dart';
import '/flutter_flow/flutter_flow_util.dart';
import '/flutter_flow/flutter_flow_widgets.dart';
import '/widget/notifikasi_card/notifikasi_card_widget.dart';
import 'dart:math';
import 'home_page_detail_widget.dart' show HomePageDetailWidget;
import 'package:flutter/material.dart';
import 'package:flutter/scheduler.dart';
import 'package:flutter_animate/flutter_animate.dart';
import 'package:google_fonts/google_fonts.dart';
import 'package:percent_indicator/percent_indicator.dart';
import 'package:provider/provider.dart';

class HomePageDetailModel extends FlutterFlowModel<HomePageDetailWidget> {
  ///  State fields for stateful widgets in this page.

  final unfocusNode = FocusNode();
  // Model for NotifikasiCard component.
  late NotifikasiCardModel notifikasiCardModel;

  @override
  void initState(BuildContext context) {
    notifikasiCardModel = createModel(context, () => NotifikasiCardModel());
  }

  @override
  void dispose() {
    notifikasiCardModel.dispose();
  }
}
