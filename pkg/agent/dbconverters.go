package agent

import (
	"fmt"
	"github.com/nimbess/nimbess-agent/pkg/etcdv3/model"
	"github.com/nimbess/nimbess-agent/pkg/network"
	log "github.com/sirupsen/logrus"
)

// NimbessIngressPortToEtcd converts the Nimbess object into an etcd representation
// This function does NOT automatically convert underlying Port objects
// Port conversion must be called separately
func NimbessIngressPortToEtcd(n *network.IngressPort) model.IngressPort {
	i := model.IngressPort{
		PipelineModule: model.PipelineModule(n.Module),
		PortName: n.PortName,
	}

	return i
}

// EtcdToNimbess converts the database representation into a new Nimbess Object
// Will automatically create underlying modules to point to for new object
func (s *NimbessAgent) EtcdToIngressPort(k model.PipelineModuleKey, i model.IngressPort) (network.IngressPort, error) {
	n := network.IngressPort{
		Module: network.Module(i.PipelineModule),
	}
	portKey := model.PortKey{
		AgentKey: k.AgentKey,
		Name: i.PortName,
	}
	// need to get port
	kv, err := s.EtcdClient.Get(s.EtcdContext, portKey)
	if err != nil {
		log.Errorf("Failure to find matching port: %s for Ingress Port: %s", i.PortName, i.Name)
		return n, err
	}

	iPort, ok := kv.Value.(network.Port)
	if !ok {
		return n, fmt.Errorf("failed to cast etcd value to a port: %v", kv.Value)
	}

	n.Port = &iPort

	return n, nil

}

// NimbessSwitchToEtcd converts the Nimbess object into an etcd representation
func NimbessSwitchToEtcd(n *network.Switch) model.Switch {
	return model.Switch(n.Module)
}

// EtcdToNimbess converts the database representation into a new Nimbess Object
// Will automatically create underlying modules to point to for new object
func (s *NimbessAgent) EtcdToSwitch(k model.PipelineModuleKey, dbSwitch model.Switch) (network.Switch, error) {
	n := network.Switch{
		Module: network.Module(dbSwitch),
	}

	// TODO need to update FIB from DB
	return n, nil
}

// NimbessToEtcd converts the database representation into a new Nimbess Pipeline
// This function does NOT automatically convert underlying module objects themselves
// Module conversion must be called separately
func NimbessMetaToEtcd(n *NimbessPipeline) model.MetaPipeline {
	m := model.MetaPipeline {
		Name: n.Name,
		Gateway: n.Gateway,
	}

	for _,mod := range n.Modules {
		m.Modules = append(m.Modules, mod.GetName())
	}

	return m
}

// EtcdToNimbess converts the database representation into a new Nimbess Pipeline
// Will automatically create underlying modules to point to for new pipeline
func (s *NimbessAgent) EtcdToMeta(m model.MetaPipeline) (NimbessPipeline, error) {
	n := NimbessPipeline{
		Name: m.Name,
		MetaKey: m.Name,
		Gateway: m.Gateway,
	}
	// TODO create underlying modules
	/**
	for _, mod := range n.Modules {
		for nimbessAgent.MetaPipelines
	}
	*/
	return n, nil
}

//NimbessToEtcd converts the database representation into a new Nimbess Pipeline
func (s *NimbessAgent) NimbessL2FIBToEtcd(fibMap map[string]*network.L2FIBEntry) model.L2FIB {
	f := model.L2FIB{}
	for k,v := range fibMap {
		f[k] = *v
	}
	return f
}

// EtcdToL2FIB converts the database representation into a new Nimbess Pipeline
func EtcdToL2FIB(f model.L2FIB) map[string]*network.L2FIBEntry {
	m := make(map[string]*network.L2FIBEntry)
	for k,v := range f {
		m[k] = &v
	}
	return m
}

//NimbessToEtcd converts the database representation into a new Nimbess Pipeline
func PipelineToEtcd(n *NimbessPipeline) (model.Pipeline, error) {
	p := model.Pipeline {
		Name: n.Name,
		MetaKey: n.MetaKey,
	}

	// if meta pipeline, store l2fib

	return p, nil
}

//EtcdToNimbess converts the database representation into a new Nimbess Pipeline
func (s *NimbessAgent) EtcdToPipeline(p model.Pipeline) (NimbessPipeline, error) {
	n := NimbessPipeline{
		Name: p.Name,
		MetaKey: p.MetaKey,
	}

	return n, nil
}