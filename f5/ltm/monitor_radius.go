// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ltm

import "e-xpert_solutions/f5-rest-client/f5"

type MonitorRadiusConfigList struct {
	Items    []MonitorRadiusConfig `json:"items"`
	Kind     string                `json:"kind"`
	SelfLink string                `json:"selflink"`
}

type MonitorRadiusConfig struct {
	Debug        string `json:"debug"`
	Destination  string `json:"destination"`
	FullPath     string `json:"fullPath"`
	Generation   int    `json:"generation"`
	Interval     int    `json:"interval"`
	Kind         string `json:"kind"`
	ManualResume string `json:"manualResume"`
	Name         string `json:"name"`
	Partition    string `json:"partition"`
	SelfLink     string `json:"selfLink"`
	TimeUntilUp  int    `json:"timeUntilUp"`
	Timeout      int    `json:"timeout"`
	UpInterval   int    `json:"upInterval"`
}

const MonitorRadiusEndpoint = "/monitor/radius"

type MonitorRadiusResource struct {
	c f5.Client
}

func (r *MonitorRadiusResource) ListAll() (*MonitorRadiusConfigList, error) {
	var list MonitorRadiusConfigList
	if err := r.c.ReadQuery(BasePath+MonitorRadiusEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

func (r *MonitorRadiusResource) Get(id string) (*MonitorRadiusConfig, error) {
	var item MonitorRadiusConfig
	if err := r.c.ReadQuery(BasePath+MonitorRadiusEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *MonitorRadiusResource) Create(item MonitorRadiusConfig) error {
	if err := r.c.ModQuery("POST", BasePath+MonitorRadiusEndpoint, item); err != nil {
		return err
	}
	return nil
}

func (r *MonitorRadiusResource) Edit(id string, item MonitorRadiusConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+MonitorRadiusEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

func (r *MonitorRadiusResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+MonitorRadiusEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
