package org.example;

//TIP To <b>Run</b> code, press <shortcut actionId="Run"/> or
// click the <icon src="AllIcons.Actions.Execute"/> icon in the gutter.
public class Main {
    public static void main(String[] args) {// TIP Press <shortcut actionId="ShowIntentionActions"/> with your caret at
                                            // the highlighted text
        try {
            Something.dosomething();
        } catch (Exception e) {
            System.out.println(e.getMessage());
        }

        var item = Item.createItem("d-1111");
        if (item == null) {
            // 何か処理
        }

    }
}