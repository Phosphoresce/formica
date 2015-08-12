package ecc

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/elb"
	"github.com/phosphoresce/formica/utils"
)

// EC2 functions that should be libraried
func resolveSg(secGroup string, client *ec2.EC2) string {
	// takes sg-xxxxxxx and client, then returns the name of the security group
	resp, err := client.DescribeSecurityGroups(nil)
	utils.HandleErr(err)
	for x := 0; x < len(resp.SecurityGroups); x++ {
		if *resp.SecurityGroups[x].GroupID == secGroup {
			secGroup = *resp.SecurityGroups[x].GroupName
			break
		}
	}
	return secGroup
}

func EntityCount() {
	regions := []string{"us-east-1", "us-west-1", "us-west-2"}
	for x := 0; x < len(regions); x++ {
		// EC2 client
		client := ec2.New(&aws.Config{Region: regions[x]})
		resp, err := client.DescribeInstances(nil)
		utils.HandleErr(err)
		// ELB client
		clientElb := elb.New(&aws.Config{Region: regions[x]})
		respElb, errElb := clientElb.DescribeLoadBalancers(nil)
		utils.HandleErr(errElb)
		fmt.Printf("Region: %v\n-----------------\n", regions[x])
		fmt.Printf("%15.15v %15.15v %10.10v %15.15v %20.20v %25.25v\n", "Instance:", "Key Pair:", "State:", "Load Balancer:", "ELB Groups:", "Security Groups:")
		for y := 0; y < len(resp.Reservations); y++ {
			for z := 0; z < len(resp.Reservations[y].Instances); z++ {
				var tag, state, key, elb, elbGroups, groups string
				if resp.Reservations[y].Instances[z].Tags[0].Value != nil {
					tag = *resp.Reservations[y].Instances[z].Tags[0].Value
				}
				if resp.Reservations[y].Instances[z].State.Name != nil {
					state = *resp.Reservations[y].Instances[z].State.Name
				}
				if resp.Reservations[y].Instances[z].KeyName != nil {
					key = *resp.Reservations[y].Instances[z].KeyName
				}
				for b := 0; b < len(respElb.LoadBalancerDescriptions); b++ {
					for c := 0; c < len(respElb.LoadBalancerDescriptions[b].Instances); c++ {
						if *resp.Reservations[y].Instances[z].InstanceID == *respElb.LoadBalancerDescriptions[b].Instances[c].InstanceID {
							if respElb.LoadBalancerDescriptions[b].LoadBalancerName != nil {
								if elb != "" {
									elb = fmt.Sprintf("%v, %.8v", elb, *respElb.LoadBalancerDescriptions[b].LoadBalancerName)
								} else {
									elb = *respElb.LoadBalancerDescriptions[b].LoadBalancerName
								}
							}
							for d := 0; d < len(respElb.LoadBalancerDescriptions[b].SecurityGroups); d++ {
								var elbGroup string
								elbGroup = *respElb.LoadBalancerDescriptions[b].SecurityGroups[d]
								elbGroup = resolveSg(elbGroup, client)
								if elbGroups != "" {
									elbGroups = fmt.Sprintf("%v, ", elbGroups)
								}
								if d < len(respElb.LoadBalancerDescriptions[b].SecurityGroups)-1 {
									elbGroups = fmt.Sprintf("%v%.8v, ", elbGroups, elbGroup)
								} else {
									elbGroups = fmt.Sprintf("%v%.8v", elbGroups, elbGroup)
								}
							}
						}
					}
				}
				for a := 0; a < len(resp.Reservations[y].Instances[z].SecurityGroups); a++ {
					var group string
					if resp.Reservations[y].Instances[z].SecurityGroups[a] != nil {
						group = *resp.Reservations[y].Instances[z].SecurityGroups[a].GroupName
						if a < len(resp.Reservations[y].Instances[z].SecurityGroups)-1 {
							groups = fmt.Sprintf("%v%.8v, ", groups, group)
						} else {
							groups = fmt.Sprintf("%v%.8v", groups, group)
						}
					}
				}
				fmt.Printf("%15.15v %15.15v %10.10v %15.15v %20.20v %25.25v\n", tag, key, state, elb, elbGroups, groups)
			}
		}
		fmt.Println()
	}
	utils.EnterTo()
	utils.ClearScreen()
}

func ServiceHealth() {
	// temp array maybe pass this in??
	regions := []string{"us-east-1", "us-west-1", "us-west-2"}
	for x := 0; x < len(regions); x++ {
		client := ec2.New(&aws.Config{Region: regions[x]})
		resp, err := client.DescribeAvailabilityZones(nil)
		utils.HandleErr(err)
		fmt.Printf("Region: %v\n-----------------\n", regions[x])
		for y := 0; y < len(resp.AvailabilityZones); y++ {
			fmt.Printf("Zone: %v... %v\n", *resp.AvailabilityZones[y].ZoneName, *resp.AvailabilityZones[y].State)
		}
		fmt.Println()
	}
	utils.EnterTo()
	utils.ClearScreen()
}

func ScheduledEvents() {
	regions := []string{"us-east-1", "us-west-1", "us-west-2"}
	events := 0
	for x := 0; x < len(regions); x++ {
		client := ec2.New(&aws.Config{Region: regions[x]})
		resp, err := client.DescribeInstanceStatus(nil)
		utils.HandleErr(err)
		for y := 0; y < len(resp.InstanceStatuses); y++ {
			for z := 0; z < len(resp.InstanceStatuses[y].Events); z++ {
				fmt.Printf("Scheduled Event Found!\nInstance: %v\nRegion: %v\nEvent Code: %v\n", resp.InstanceStatuses[y].InstanceID, regions[x], resp.InstanceStatuses[y].Events[z].Code)
				fmt.Println()
				events++
			}
		}
	}
	fmt.Printf("%v events found.\n\n", events)
	utils.EnterTo()
	utils.ClearScreen()
}

// -- end -- //
