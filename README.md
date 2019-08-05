# 基于Golang开发的flutter生产工具

## flutter资源配置生成工具
### 自动化标识
```
flutter:

  # The following line ensures that the Material Icons font is
  # included with your application, so that you can use the icons in
  # the material Icons class.
  uses-material-design: true

  # To add assets to your application, add an assets section, like this:
  assets:
  ## <<assets begin>>[起始标识]
    - images/btn_denglu.png
    - images/common/btn_chehui.png
  ## <<assets end>>[结束标识]
```