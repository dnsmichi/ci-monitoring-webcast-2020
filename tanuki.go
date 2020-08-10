package main

import (
  "fmt"
)

func Color(colorString string) func(...interface{}) string {
  sprint := func(args ...interface{}) string {
    return fmt.Sprintf(colorString,
      fmt.Sprint(args...))
  }
  return sprint
}

var (
  earNoseColor = Color("\033[38;5;88m%s\033[0m");
  eyesColor = Color("\033[38;5;208m%s\033[0m");
  cheeksColor = Color("\033[38;5;202m%s\033[0m");
)

func GetTanuki(mode bool) string {
  var out string;

  //Borrowed from https://gitlab.com/gitlab-org/gitlab-foss/-/raw/4086e41fbbcc9d3fe21adf2fb404d778b699d9d7/app/assets/javascripts/console_swag.js
  //Note: Avoid multi line strings in Golang, escaping backticks is a pain.
  out += earNoseColor("           +                        +") + "\n";
  out += earNoseColor("          :s:                      :s:") + "\n";
  out += earNoseColor("         .oso'                    'oso.") + "\n";
  out += earNoseColor("         +sss+                    +sss+") + "\n";
  out += earNoseColor("        :sssss-                  -sssss:") + "\n";
  out += earNoseColor("       'ossssso.                .ossssso'") + "\n";
  out += earNoseColor("       +sssssss+                +sssssss+") + "\n";
  out += earNoseColor("      -ooooooooo-++++++++++++++-ooooooooo-") + "\n";
  out += cheeksColor("     ':/") + eyesColor("+++++++++") + earNoseColor("osssssssssssso") + eyesColor("+++++++++") + cheeksColor("/:'") + "\n";
  out += cheeksColor("     -///") + eyesColor("+++++++++") + earNoseColor("cssssssssssss") + eyesColor("+++++++++") + cheeksColor("///-") + "\n";
  out += cheeksColor("    .//////") + eyesColor("+++++++") + earNoseColor("cosssssssssso") + eyesColor("+++++++") + cheeksColor("//////.") + "\n";
  out += cheeksColor("    :///////") + eyesColor("+++++++") + earNoseColor("osssssssso") + eyesColor("+++++++") + cheeksColor("///////:") + "\n";
  out += cheeksColor("     .:///////") + eyesColor("++++++") + earNoseColor("ssssssss") + eyesColor("++++++") + cheeksColor("///////:.'") + "\n";
  out += cheeksColor("       '-://///") + eyesColor("+++++") + earNoseColor("osssssso") + eyesColor("+++++") + cheeksColor("/////:-'") + "\n";
  out += cheeksColor("          '-:////") + eyesColor("++++") + earNoseColor("osssso") + eyesColor("++++") + cheeksColor("////:-'") + "\n";
  out += cheeksColor("             .-:///") + eyesColor("++") + earNoseColor("osssso") + eyesColor("++") + cheeksColor("///:-.") + "\n";
  out += cheeksColor("               '.://") + eyesColor("++") + earNoseColor("cosso") + eyesColor("++") + cheeksColor("//:.'") + "\n";
  out += cheeksColor("                  '-:/") + eyesColor("+") + earNoseColor("coo") + eyesColor("+") + cheeksColor("/:-'") + "\n";
  out += cheeksColor("                     '-") + eyesColor("++") + cheeksColor("-'") + "\n";
  out += "\n\n" + cheeksColor("Tanuki has been summoned...") + "\n";

  return out;
}
