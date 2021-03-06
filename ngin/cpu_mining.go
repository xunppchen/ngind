// Copyright 2015 The go-ethereum Authors
// This file is part of Ngin.
//
// Ngin is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// Ngin is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with Ngin. If not, see <http://www.gnu.org/licenses/>.

// +build !opencl

package ngin

import (
	"errors"
	"fmt"

	"github.com/NginProject/ngind/logger"
	"github.com/NginProject/ngind/logger/glog"
)

func (s *Ngin) StartMining(threads int, gpus string) error {
	eb, err := s.Coinbase()
	if err != nil {
		err = fmt.Errorf("Cannot start mining without coinbase address: %v", err)
		glog.V(logger.Error).Infoln(err)
		return err
	}

	if gpus != "" {
		return errors.New("GPU mining not supported yet.")
	}

	// CPU mining
	go s.miner.Start(eb, threads)
	return nil
}
