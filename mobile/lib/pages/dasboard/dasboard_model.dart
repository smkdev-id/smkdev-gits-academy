import '/flutter_flow/flutter_flow_charts.dart';
import '/flutter_flow/flutter_flow_icon_button.dart';
import '/flutter_flow/flutter_flow_theme.dart';
import '/flutter_flow/flutter_flow_util.dart';
import '/flutter_flow/flutter_flow_widgets.dart';
import '/widget/history_paymets_card/history_paymets_card_widget.dart';
import '/flutter_flow/random_data_util.dart' as random_data;
import 'dasboard_widget.dart' show DasboardWidget;
import 'package:flutter/material.dart';
import 'package:google_fonts/google_fonts.dart';
import 'package:provider/provider.dart';

class DasboardModel extends FlutterFlowModel<DasboardWidget> {
  ///  State fields for stateful widgets in this page.

  final unfocusNode = FocusNode();
  // Model for HistoryPaymetsCard component.
  late HistoryPaymetsCardModel historyPaymetsCardModel1;
  // Model for HistoryPaymetsCard component.
  late HistoryPaymetsCardModel historyPaymetsCardModel2;
  // Model for HistoryPaymetsCard component.
  late HistoryPaymetsCardModel historyPaymetsCardModel3;
  // Model for HistoryPaymetsCard component.
  late HistoryPaymetsCardModel historyPaymetsCardModel4;

  @override
  void initState(BuildContext context) {
    historyPaymetsCardModel1 =
        createModel(context, () => HistoryPaymetsCardModel());
    historyPaymetsCardModel2 =
        createModel(context, () => HistoryPaymetsCardModel());
    historyPaymetsCardModel3 =
        createModel(context, () => HistoryPaymetsCardModel());
    historyPaymetsCardModel4 =
        createModel(context, () => HistoryPaymetsCardModel());
  }

  @override
  void dispose() {
    historyPaymetsCardModel1.dispose();
    historyPaymetsCardModel2.dispose();
    historyPaymetsCardModel3.dispose();
    historyPaymetsCardModel4.dispose();
  }
}
