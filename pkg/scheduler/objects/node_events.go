/*
 Licensed to the Apache Software Foundation (ASF) under one
 or more contributor license agreements.  See the NOTICE file
 distributed with this work for additional information
 regarding copyright ownership.  The ASF licenses this file
 to you under the Apache License, Version 2.0 (the
 "License"); you may not use this file except in compliance
 with the License.  You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

package objects

import (
	"github.com/apache/yunikorn-core/pkg/common"
	"github.com/apache/yunikorn-core/pkg/common/resources"
	"github.com/apache/yunikorn-core/pkg/events"
	"github.com/apache/yunikorn-scheduler-interface/lib/go/si"
)

type nodeEvents struct {
	enabled     bool
	eventSystem events.EventSystem
	node        *Node
}

func (n *nodeEvents) sendNodeAddedEvent() {
	if !n.enabled {
		return
	}

	event := events.CreateNodeEventRecord(n.node.NodeID, common.Empty, common.Empty, si.EventRecord_ADD,
		si.EventRecord_DETAILS_NONE, n.node.GetCapacity())
	n.eventSystem.AddEvent(event)
}

func (n *nodeEvents) sendNodeRemovedEvent() {
	if !n.enabled {
		return
	}

	event := events.CreateNodeEventRecord(n.node.NodeID, common.Empty, common.Empty, si.EventRecord_REMOVE,
		si.EventRecord_NODE_DECOMISSION, nil)
	n.eventSystem.AddEvent(event)
}

func (n *nodeEvents) sendAllocationAddedEvent(allocID string, res *resources.Resource) {
	if !n.enabled {
		return
	}

	event := events.CreateNodeEventRecord(n.node.NodeID, common.Empty, allocID, si.EventRecord_ADD,
		si.EventRecord_NODE_ALLOC, res)
	n.eventSystem.AddEvent(event)
}

func (n *nodeEvents) sendAllocationRemovedEvent(allocID string, res *resources.Resource) {
	if !n.enabled {
		return
	}

	event := events.CreateNodeEventRecord(n.node.NodeID, common.Empty, allocID, si.EventRecord_REMOVE,
		si.EventRecord_NODE_ALLOC, res)
	n.eventSystem.AddEvent(event)
}

func (n *nodeEvents) sendNodeReadyChangedEvent(ready bool) {
	if !n.enabled {
		return
	}

	reason := ""
	if ready {
		reason = "ready: true"
	} else {
		reason = "ready: false"
	}

	event := events.CreateNodeEventRecord(n.node.NodeID, reason, common.Empty, si.EventRecord_SET,
		si.EventRecord_NODE_READY, nil)
	n.eventSystem.AddEvent(event)
}

func (n *nodeEvents) sendNodeSchedulableChangedEvent(ready bool) {
	if !n.enabled {
		return
	}

	var reason string
	if ready {
		reason = "schedulable: true"
	} else {
		reason = "schedulable: false"
	}

	event := events.CreateNodeEventRecord(n.node.NodeID, reason, common.Empty, si.EventRecord_SET,
		si.EventRecord_NODE_SCHEDULABLE, nil)
	n.eventSystem.AddEvent(event)
}

func (n *nodeEvents) sendNodeCapacityChangedEvent() {
	if !n.enabled {
		return
	}

	event := events.CreateNodeEventRecord(n.node.NodeID, common.Empty, common.Empty, si.EventRecord_SET,
		si.EventRecord_NODE_CAPACITY, n.node.totalResource)
	n.eventSystem.AddEvent(event)
}

func (n *nodeEvents) sendNodeOccupiedResourceChangedEvent() {
	if !n.enabled {
		return
	}

	event := events.CreateNodeEventRecord(n.node.NodeID, common.Empty, common.Empty, si.EventRecord_SET,
		si.EventRecord_NODE_OCCUPIED, n.node.occupiedResource)
	n.eventSystem.AddEvent(event)
}

func (n *nodeEvents) sendReservedEvent(res *resources.Resource, askID string) {
	if !n.enabled {
		return
	}

	event := events.CreateNodeEventRecord(n.node.NodeID, common.Empty, askID, si.EventRecord_ADD,
		si.EventRecord_NODE_RESERVATION, res)
	n.eventSystem.AddEvent(event)
}

func (n *nodeEvents) sendUnreservedEvent(res *resources.Resource, askID string) {
	if !n.enabled {
		return
	}

	event := events.CreateNodeEventRecord(n.node.NodeID, common.Empty, askID, si.EventRecord_REMOVE,
		si.EventRecord_NODE_RESERVATION, res)
	n.eventSystem.AddEvent(event)
}

func newNodeEvents(node *Node, evt events.EventSystem) *nodeEvents {
	return &nodeEvents{
		eventSystem: evt,
		enabled:     evt != nil,
		node:        node,
	}
}
