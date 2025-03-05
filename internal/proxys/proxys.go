package proxys

import "sync"

type IProxy struct {
	ips  []string
	lock sync.Mutex
}

func NewIProxyIP(ips []string) IProxy {
	return IProxy{
		ips: ips,
	}
}

func (p *IProxy) GetIPS() int {
	return len(p.ips)
}

func (p *IProxy) GetProxyIP() string {
	if p == nil {
		return ""
	}

	p.lock.Lock()
	defer p.lock.Unlock()

	if len(p.ips) == 0 {
		return ""
	}

	proxyIp := p.ips[0]
	p.ips = append(p.ips[1:], proxyIp)
	return proxyIp
}

func (p *IProxy) RemoveProxyIP(proxyIp string) {
	if p == nil {
		return
	}

	p.lock.Lock()
	defer p.lock.Unlock()

	for i, ip := range p.ips {
		if ip == proxyIp {
			// Remove the proxy IP from the slice
			p.ips = append(p.ips[:i], p.ips[i+1:]...)
			break
		}
	}
}
