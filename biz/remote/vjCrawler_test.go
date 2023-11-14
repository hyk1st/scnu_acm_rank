package remote

import "testing"

func TestLogin(t *testing.T) {

	vj := VjCrawler{
		userName: "3553928717@qq.com",
		passWord: "wsy16675060764",
		cookie:   "",
	}

	_, err := vj.Login()
	if err != nil {
		t.Fatal(err)
	}
}

func TestCheckLoginStatus(t *testing.T) {
	vj := VjCrawler{
		userName: "3553928717@qq.com",
		passWord: "wsy16675060764",
		cookie:   "_ga=GA1.1.713543795.1666859533; JSESSlONID=1S22VY2F5CKKJLDSY63UE8LDNJAXZTSH; _ga_374JLX1715=GS1.1.1699940732.9.1.1699941250.54.0.0; JSESSIONID=53B09484476B0D27FE315B8E9CB598E0; Jax.Q=123123213|NDTRGCKLIYA41KP0PWTVSBCTCW07V3",
	}
	if f, err := vj.checkLoginStatus(); !f || err != nil {
		t.Fatal(f)
	}
}
