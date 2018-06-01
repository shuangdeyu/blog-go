package controller

import (
	"strconv"
	"math"
)

const (
	full_tag_open = "<ul class=\"pagination\">" // 分页总模块的起始标签
	full_tag_close = "</ul>" // 分页总模块的结束标签
	first_tag_open = "<li>" // 第一个链接的起始标签
	first_tag_close = "</li>" // 第一个链接的结束标签
	prev_tag_open = "<li>" // 上一页链接的起始标签
	prev_tag_close = "</li>" // 上一页链接的结束标签
	next_tag_open = "<li>" // 下一页链接的起始标签
	next_tag_close = "</li>" // 下一页链接的结束标签
	cur_tag_open = "<li class=\"active\"><a>" // 当前页链接的起始标签
	cur_tag_close = "</a></li>" // 当前页链接的结束标签
	last_tag_open = "<li>" // 最后一个链接的起始标签
	last_tag_close = "</li>" // 最后一个链接的结束标签
	num_tag_open = "<li>" // 数字链接的起始标签
	num_tag_close = "</li>" // 数字链接的结束标签
	attributes = "class=\"myclass\"" // 给所有<a>标签加上class
	first_link = "首页" // 第一个链接的名称
	next_link = "下一页" // 下一页的名称
	prev_link = "上一页" // 上一页的名称
	last_link = "末页" // 最后一个链接的名称
)

func CreatePage(total int, page_size int, cur_page string, base_url string) (string, error) {
	page, _ := strconv.Atoi(cur_page)
	page_str := ""
	// 总页数
	num_page := int(math.Ceil(float64(total) / float64(page_size)))
	// 只有一页
	if num_page <= 1 {
		return page_str, nil
	}

	page_str = full_tag_open // 分页模块开始
	li_str := ""
	if num_page <= 3 {
		// 3 页之内处理
		for i :=  0;i < num_page;i++ {
			p := strconv.Itoa(i+1)
			if i == page {
				li_str += cur_tag_open + p + cur_tag_close
			} else {
				li_str += num_tag_open + "<a href=\"" + base_url + "/" + strconv.Itoa(i) + "\" " + attributes + " >" + p + "</a>" + num_tag_close
			}
		}
		if page < (num_page-1) {
			li_str += next_tag_open + "<a href=\"" + base_url + "/" + strconv.Itoa(page+1) + "\" " + attributes + " >" + next_link + "</a>" + next_tag_close
		}
		if page > 0 {
			li_str = prev_tag_open + "<a href=\"" + base_url + "/" + strconv.Itoa(page-1) + "\" " + attributes + " >" + prev_link + "</a>" + prev_tag_close + li_str
		}
	} else {
		// 3 页之外处理
		if page == 0 {
			li_str += cur_tag_open + "1" + cur_tag_close
			li_str += num_tag_open + "<a href=\"" + base_url + "/" + "1" + "\" " + attributes + " >" + "2" + "</a>" + num_tag_close
			li_str += num_tag_open + "<a href=\"" + base_url + "/" + "2" + "\" " + attributes + " >" + "3" + "</a>" + num_tag_close
		} else if page == 1 {
			li_str += num_tag_open + "<a href=\"" + base_url + "/" + "0" + "\" " + attributes + " >" + "1" + "</a>" + num_tag_close
			li_str += cur_tag_open + "2" + cur_tag_close
			li_str += num_tag_open + "<a href=\"" + base_url + "/" + "2" + "\" " + attributes + " >" + "3" + "</a>" + num_tag_close
			li_str += num_tag_open + "<a href=\"" + base_url + "/" + "3" + "\" " + attributes + " >" + "4" + "</a>" + num_tag_close
		} else if page == (num_page-1) {
			li_str += num_tag_open + "<a href=\"" + base_url + "/" + strconv.Itoa(num_page-3) + "\" " + attributes + " >" + strconv.Itoa(num_page-2) + "</a>" + num_tag_close
			li_str += num_tag_open + "<a href=\"" + base_url + "/" + strconv.Itoa(num_page-2) + "\" " + attributes + " >" + strconv.Itoa(num_page-1) + "</a>" + num_tag_close
			li_str += cur_tag_open + strconv.Itoa(num_page) + cur_tag_close
		} else if page == (num_page-2) {
			li_str += num_tag_open + "<a href=\"" + base_url + "/" + strconv.Itoa(num_page-4) + "\" " + attributes + " >" + strconv.Itoa(num_page-3) + "</a>" + num_tag_close
			li_str += num_tag_open + "<a href=\"" + base_url + "/" + strconv.Itoa(num_page-3) + "\" " + attributes + " >" + strconv.Itoa(num_page-2) + "</a>" + num_tag_close
			li_str += cur_tag_open + strconv.Itoa(num_page-1) + cur_tag_close
			li_str += num_tag_open + "<a href=\"" + base_url + "/" + strconv.Itoa(num_page-1) + "\" " + attributes + " >" + strconv.Itoa(num_page) + "</a>" + num_tag_close
		} else {
			li_str += num_tag_open + "<a href=\"" + base_url + "/" + strconv.Itoa(page-2) + "\" " + attributes + " >" + strconv.Itoa(page-1) + "</a>" + num_tag_close
			li_str += num_tag_open + "<a href=\"" + base_url + "/" + strconv.Itoa(page-1) + "\" " + attributes + " >" + strconv.Itoa(page) + "</a>" + num_tag_close
			li_str += cur_tag_open + strconv.Itoa(page+1) + cur_tag_close
			li_str += num_tag_open + "<a href=\"" + base_url + "/" + strconv.Itoa(page+1) + "\" " + attributes + " >" + strconv.Itoa(page+2) + "</a>" + num_tag_close
			li_str += num_tag_open + "<a href=\"" + base_url + "/" + strconv.Itoa(page+2) + "\" " + attributes + " >" + strconv.Itoa(page+3) + "</a>" + num_tag_close
		}
		if page < (num_page-1) {
			li_str += next_tag_open + "<a href=\"" + base_url + "/" + strconv.Itoa(page+1) + "\" " + attributes + " >" + next_link + "</a>" + next_tag_close
		}
		if page > 0 {
			li_str = prev_tag_open + "<a href=\"" + base_url + "/" + strconv.Itoa(page-1) + "\" " + attributes + " >" + prev_link + "</a>" + prev_tag_close + li_str
		}
		if page < (num_page - 2) {
			li_str += last_tag_open + "<a href=\"" + base_url + "/" + strconv.Itoa(num_page - 1) + "\" " + attributes + " >" + last_link + "</a>" + last_tag_close
		}
		if page > 1 {
			li_str = first_tag_open + "<a href=\"" + base_url + "/" + "0" + "\" " + attributes + " >" + first_link + "</a>" + first_tag_close + li_str
		}
	}

	page_str += li_str + full_tag_close // 分页模块结束
	return page_str, nil
}