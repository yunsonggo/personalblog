import VXETable from 'vxe-table'

export default function vxeTableAddFormats() {
  VXETable.formats.add('formatMenu', ({ cellValue }, menuList) => {
    var res = ""
    menuList.forEach((menu) => {
      if (cellValue == menu.Id) {
        res = menu.title
      }
    })
    return res
  })
  VXETable.formats.add('formatCategory', ({ cellValue }, categoryList) => {
    var res = ""
    if (cellValue == 0) {
      res = "无"
    } else {
      categoryList.forEach((cate) => {
        if (cellValue == cate.Id) {
          res = cate.title
        }
      })
    }
    return res
  })
  VXETable.formats.add('formatOriginal', ({ cellValue }) => {
    var res = ""
    if (cellValue == 0) {
      res = "原创"
    } else {
      res = "引用"
    }
    return res
  })
}



