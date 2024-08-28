import '/flutter_flow/flutter_flow_theme.dart';
import '/flutter_flow/flutter_flow_util.dart';
import '/flutter_flow/flutter_flow_widgets.dart';
import '/widget/notifikasi_card/notifikasi_card_widget.dart';
import 'notifikasi_page_widget.dart' show NotifikasiPageWidget;
import 'package:flutter/material.dart';
import 'package:google_fonts/google_fonts.dart';
import 'package:provider/provider.dart';

class NotifikasiPageModel extends FlutterFlowModel<NotifikasiPageWidget> {
  ///  State fields for stateful widgets in this page.

  // Model for NotifikasiCard component.
  late NotifikasiCardModel notifikasiCardModel1;
  // Model for NotifikasiCard component.
  late NotifikasiCardModel notifikasiCardModel2;
  // Model for NotifikasiCard component.
  late NotifikasiCardModel notifikasiCardModel3;

  @override
  void initState(BuildContext context) {
    notifikasiCardModel1 = createModel(context, () => NotifikasiCardModel());
    notifikasiCardModel2 = createModel(context, () => NotifikasiCardModel());
    notifikasiCardModel3 = createModel(context, () => NotifikasiCardModel());
  }

  @override
  void dispose() {
    notifikasiCardModel1.dispose();
    notifikasiCardModel2.dispose();
    notifikasiCardModel3.dispose();
  }
}
