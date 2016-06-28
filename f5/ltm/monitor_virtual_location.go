// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ltm

import "e-xpert_solutions/f5-rest-client/f5"

type MonitorVirtualLocationConfigList struct {
	Items    []MonitorVirtualLocationConfig `json:"items"`
	Kind     string                         `json:"kind"`
	SelfLink string                         `json:"selflink"`
}

type MonitorVirtualLocationConfig struct {
	Debug       string `json:"debug"`
	FullPath    string `json:"fullPath"`
	Generation  int    `json:"generation"`
	Interval    int    `json:"interval"`
	Kind        string `json:"kind"`
	Name        string `json:"name"`
	Partition   string `json:"partition"`
	SelfLink    string `json:"selfLink"`
	TimeUntilUp int    `json:"timeUntilUp"`
	Timeout     int    `json:"timeout"`
	UpInterval  int    `json:"upInterval"`
}

const MonitorVirtualLocationEndpoint = "/monitor/virtual-location"

type MonitorVirtualLocationResource struct {
	c f5.Client
}

func (r *MonitorVirtualLocationResource) ListAll() (*MonitorVirtualLocationConfigList, error) {
	var list MonitorVirtualLocationConfigList
	if err := r.c.ReadQuery(BasePath+MonitorVirtualLocationEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

func (r *MonitorVirtualLocationResource) Get(id string) (*MonitorVirtualLocationConfig, error) {
	var item MonitorVirtualLocationConfig
	if err := r.c.ReadQuery(BasePath+MonitorVirtualLocationEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *MonitorVirtualLocationResource) Create(item MonitorVirtualLocationConfig) error {
	if err := r.c.ModQuery("POST", BasePath+MonitorVirtualLocationEndpoint, item); err != nil {
		return err
	}
	return nil
}

func (r *MonitorVirtualLocationResource) Edit(id string, item MonitorVirtualLocationConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+MonitorVirtualLocationEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

func (r *MonitorVirtualLocationResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+MonitorVirtualLocationEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
