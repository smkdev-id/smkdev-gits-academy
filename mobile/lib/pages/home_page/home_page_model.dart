import '/flutter_flow/flutter_flow_animations.dart';
import '/flutter_flow/flutter_flow_theme.dart';
import '/flutter_flow/flutter_flow_util.dart';
import '/flutter_flow/flutter_flow_widgets.dart';
import '/widget/income_card_home_page/income_card_home_page_widget.dart';
import '/widget/notifikasi_card/notifikasi_card_widget.dart';
import 'dart:math';
import 'home_page_widget.dart' show HomePageWidget;
import 'package:flutter/material.dart';
import 'package:flutter/scheduler.dart';
import 'package:flutter_animate/flutter_animate.dart';
import 'package:google_fonts/google_fonts.dart';
import 'package:provider/provider.dart';

class HomePageModel extends FlutterFlowModel<HomePageWidget> {
  ///  State fields for stateful widgets in this page.

  final unfocusNode = FocusNode();
  // Model for incomeCardHomePage component.
  late IncomeCardHomePageModel incomeCardHomePageModel1;
  // Model for incomeCardHomePage component.
  late IncomeCardHomePageModel incomeCardHomePageModel2;
  // Model for incomeCardHomePage component.
  late IncomeCardHomePageModel incomeCardHomePageModel3;
  // Model for incomeCardHomePage component.
  late IncomeCardHomePageModel incomeCardHomePageModel4;
  // Model for NotifikasiCard component.
  late NotifikasiCardModel notifikasiCardModel;

  @override
  void initState(BuildContext context) {
    incomeCardHomePageModel1 =
        createModel(context, () => IncomeCardHomePageModel());
    incomeCardHomePageModel2 =
        createModel(context, () => IncomeCardHomePageModel());
    incomeCardHomePageModel3 =
        createModel(context, () => IncomeCardHomePageModel());
    incomeCardHomePageModel4 =
        createModel(context, () => IncomeCardHomePageModel());
    notifikasiCardModel = createModel(context, () => NotifikasiCardModel());
  }

  @override
  void dispose() {
    incomeCardHomePageModel1.dispose();
    incomeCardHomePageModel2.dispose();
    incomeCardHomePageModel3.dispose();
    incomeCardHomePageModel4.dispose();
    notifikasiCardModel.dispose();
  }
}
