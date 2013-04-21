// NOTE: generated automatically by blendef for Blender v266.

package block

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

// ParseBody parses the block body and stores it in blk.Body. It is safe to call
// ParseBody multiple times on the same block.
func (blk *Block) ParseBody(order binary.ByteOrder, dna *DNA) (err error) {
	// Get block body reader.
	r, ok := blk.Body.(io.Reader)
	if !ok {
		// Body has already been parsed.
		return nil
	}

	index := blk.Hdr.SDNAIndex
	if index == 0 {
		// Parse based on block code.
		switch blk.Hdr.Code {
		case CodeDATA:
			blk.Body, err = ioutil.ReadAll(r)
			if err != nil {
				return err
			}
		case CodeDNA1:
			blk.Body, err = ParseDNA(r, order)
			if err != nil {
				return err
			}
		case CodeREND, CodeTEST:
			/// TODO: implement specific block body parsing for REND and TEST.
			blk.Body, err = ioutil.ReadAll(r)
			if err != nil {
				return err
			}
		default:
			return fmt.Errorf("Block.ParseBody: parsing of %q not yet implemented.", blk.Hdr.Code)
		}
	} else {
		// Parse based on SDNA index.
		typ := dna.Structs[index].Type
		switch typ {
		case "ARegion":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*ARegion, blk.Hdr.Count)
				for i := range bodies {
					body := new(ARegion)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(ARegion)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "AnimData":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*AnimData, blk.Hdr.Count)
				for i := range bodies {
					body := new(AnimData)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(AnimData)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "AnimMapPair":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*AnimMapPair, blk.Hdr.Count)
				for i := range bodies {
					body := new(AnimMapPair)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(AnimMapPair)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "AnimMapper":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*AnimMapper, blk.Hdr.Count)
				for i := range bodies {
					body := new(AnimMapper)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(AnimMapper)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "AnimOverride":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*AnimOverride, blk.Hdr.Count)
				for i := range bodies {
					body := new(AnimOverride)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(AnimOverride)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "ArmatureModifierData":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*ArmatureModifierData, blk.Hdr.Count)
				for i := range bodies {
					body := new(ArmatureModifierData)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(ArmatureModifierData)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "ArrayModifierData":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*ArrayModifierData, blk.Hdr.Count)
				for i := range bodies {
					body := new(ArrayModifierData)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(ArrayModifierData)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "AudioData":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*AudioData, blk.Hdr.Count)
				for i := range bodies {
					body := new(AudioData)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(AudioData)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "AviCodecData":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*AviCodecData, blk.Hdr.Count)
				for i := range bodies {
					body := new(AviCodecData)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(AviCodecData)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "BGpic":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BGpic, blk.Hdr.Count)
				for i := range bodies {
					body := new(BGpic)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BGpic)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "BMeshModifierData":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BMeshModifierData, blk.Hdr.Count)
				for i := range bodies {
					body := new(BMeshModifierData)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BMeshModifierData)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "BPoint":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BPoint, blk.Hdr.Count)
				for i := range bodies {
					body := new(BPoint)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BPoint)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "Base":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*Base, blk.Hdr.Count)
				for i := range bodies {
					body := new(Base)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(Base)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "BevelModifierData":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BevelModifierData, blk.Hdr.Count)
				for i := range bodies {
					body := new(BevelModifierData)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BevelModifierData)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "BezTriple":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BezTriple, blk.Hdr.Count)
				for i := range bodies {
					body := new(BezTriple)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BezTriple)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "BoidData":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BoidData, blk.Hdr.Count)
				for i := range bodies {
					body := new(BoidData)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BoidData)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "BoidParticle":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BoidParticle, blk.Hdr.Count)
				for i := range bodies {
					body := new(BoidParticle)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BoidParticle)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "BoidRule":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BoidRule, blk.Hdr.Count)
				for i := range bodies {
					body := new(BoidRule)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BoidRule)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "BoidRuleAverageSpeed":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BoidRuleAverageSpeed, blk.Hdr.Count)
				for i := range bodies {
					body := new(BoidRuleAverageSpeed)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BoidRuleAverageSpeed)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "BoidRuleAvoidCollision":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BoidRuleAvoidCollision, blk.Hdr.Count)
				for i := range bodies {
					body := new(BoidRuleAvoidCollision)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BoidRuleAvoidCollision)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "BoidRuleFight":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BoidRuleFight, blk.Hdr.Count)
				for i := range bodies {
					body := new(BoidRuleFight)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BoidRuleFight)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "BoidRuleFollowLeader":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BoidRuleFollowLeader, blk.Hdr.Count)
				for i := range bodies {
					body := new(BoidRuleFollowLeader)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BoidRuleFollowLeader)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "BoidRuleGoalAvoid":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BoidRuleGoalAvoid, blk.Hdr.Count)
				for i := range bodies {
					body := new(BoidRuleGoalAvoid)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BoidRuleGoalAvoid)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "BoidSettings":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BoidSettings, blk.Hdr.Count)
				for i := range bodies {
					body := new(BoidSettings)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BoidSettings)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "BoidState":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BoidState, blk.Hdr.Count)
				for i := range bodies {
					body := new(BoidState)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BoidState)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "Bone":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*Bone, blk.Hdr.Count)
				for i := range bodies {
					body := new(Bone)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(Bone)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "BooleanModifierData":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BooleanModifierData, blk.Hdr.Count)
				for i := range bodies {
					body := new(BooleanModifierData)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BooleanModifierData)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "BoundBox":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BoundBox, blk.Hdr.Count)
				for i := range bodies {
					body := new(BoundBox)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BoundBox)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "BrightContrastModifierData":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BrightContrastModifierData, blk.Hdr.Count)
				for i := range bodies {
					body := new(BrightContrastModifierData)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BrightContrastModifierData)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "Brush":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*Brush, blk.Hdr.Count)
				for i := range bodies {
					body := new(Brush)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(Brush)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "BrushClone":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BrushClone, blk.Hdr.Count)
				for i := range bodies {
					body := new(BrushClone)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BrushClone)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "BuildEff":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BuildEff, blk.Hdr.Count)
				for i := range bodies {
					body := new(BuildEff)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BuildEff)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "BuildModifierData":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BuildModifierData, blk.Hdr.Count)
				for i := range bodies {
					body := new(BuildModifierData)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BuildModifierData)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "BulletSoftBody":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BulletSoftBody, blk.Hdr.Count)
				for i := range bodies {
					body := new(BulletSoftBody)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BulletSoftBody)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "CBData":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*CBData, blk.Hdr.Count)
				for i := range bodies {
					body := new(CBData)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(CBData)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "Camera":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*Camera, blk.Hdr.Count)
				for i := range bodies {
					body := new(Camera)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(Camera)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "CastModifierData":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*CastModifierData, blk.Hdr.Count)
				for i := range bodies {
					body := new(CastModifierData)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(CastModifierData)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "ChannelDriver":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*ChannelDriver, blk.Hdr.Count)
				for i := range bodies {
					body := new(ChannelDriver)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(ChannelDriver)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "CharInfo":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*CharInfo, blk.Hdr.Count)
				for i := range bodies {
					body := new(CharInfo)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(CharInfo)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "ChildParticle":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*ChildParticle, blk.Hdr.Count)
				for i := range bodies {
					body := new(ChildParticle)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(ChildParticle)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "ClothCollSettings":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*ClothCollSettings, blk.Hdr.Count)
				for i := range bodies {
					body := new(ClothCollSettings)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(ClothCollSettings)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "ClothModifierData":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*ClothModifierData, blk.Hdr.Count)
				for i := range bodies {
					body := new(ClothModifierData)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(ClothModifierData)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "ClothSimSettings":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*ClothSimSettings, blk.Hdr.Count)
				for i := range bodies {
					body := new(ClothSimSettings)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(ClothSimSettings)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "CollisionModifierData":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*CollisionModifierData, blk.Hdr.Count)
				for i := range bodies {
					body := new(CollisionModifierData)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(CollisionModifierData)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "ColorBalanceModifierData":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*ColorBalanceModifierData, blk.Hdr.Count)
				for i := range bodies {
					body := new(ColorBalanceModifierData)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(ColorBalanceModifierData)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "ColorBand":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*ColorBand, blk.Hdr.Count)
				for i := range bodies {
					body := new(ColorBand)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(ColorBand)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "ColorCorrectionData":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*ColorCorrectionData, blk.Hdr.Count)
				for i := range bodies {
					body := new(ColorCorrectionData)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(ColorCorrectionData)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "ColorManagedColorspaceSettings":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*ColorManagedColorspaceSettings, blk.Hdr.Count)
				for i := range bodies {
					body := new(ColorManagedColorspaceSettings)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(ColorManagedColorspaceSettings)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "ColorManagedDisplaySettings":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*ColorManagedDisplaySettings, blk.Hdr.Count)
				for i := range bodies {
					body := new(ColorManagedDisplaySettings)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(ColorManagedDisplaySettings)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "ColorManagedViewSettings":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*ColorManagedViewSettings, blk.Hdr.Count)
				for i := range bodies {
					body := new(ColorManagedViewSettings)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(ColorManagedViewSettings)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "ColorMapping":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*ColorMapping, blk.Hdr.Count)
				for i := range bodies {
					body := new(ColorMapping)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(ColorMapping)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "ConsoleLine":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*ConsoleLine, blk.Hdr.Count)
				for i := range bodies {
					body := new(ConsoleLine)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(ConsoleLine)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "Curve":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*Curve, blk.Hdr.Count)
				for i := range bodies {
					body := new(Curve)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(Curve)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "CurveMap":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*CurveMap, blk.Hdr.Count)
				for i := range bodies {
					body := new(CurveMap)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(CurveMap)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "CurveMapPoint":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*CurveMapPoint, blk.Hdr.Count)
				for i := range bodies {
					body := new(CurveMapPoint)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(CurveMapPoint)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "CurveMapping":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*CurveMapping, blk.Hdr.Count)
				for i := range bodies {
					body := new(CurveMapping)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(CurveMapping)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "CurveModifierData":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*CurveModifierData, blk.Hdr.Count)
				for i := range bodies {
					body := new(CurveModifierData)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(CurveModifierData)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "CurvesModifierData":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*CurvesModifierData, blk.Hdr.Count)
				for i := range bodies {
					body := new(CurvesModifierData)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(CurvesModifierData)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "CustomData":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*CustomData, blk.Hdr.Count)
				for i := range bodies {
					body := new(CustomData)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(CustomData)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "CustomDataExternal":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*CustomDataExternal, blk.Hdr.Count)
				for i := range bodies {
					body := new(CustomDataExternal)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(CustomDataExternal)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "CustomDataLayer":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*CustomDataLayer, blk.Hdr.Count)
				for i := range bodies {
					body := new(CustomDataLayer)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(CustomDataLayer)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "DecimateModifierData":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*DecimateModifierData, blk.Hdr.Count)
				for i := range bodies {
					body := new(DecimateModifierData)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(DecimateModifierData)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "DisplaceModifierData":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*DisplaceModifierData, blk.Hdr.Count)
				for i := range bodies {
					body := new(DisplaceModifierData)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(DisplaceModifierData)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "DriverTarget":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*DriverTarget, blk.Hdr.Count)
				for i := range bodies {
					body := new(DriverTarget)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(DriverTarget)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "DriverVar":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*DriverVar, blk.Hdr.Count)
				for i := range bodies {
					body := new(DriverVar)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(DriverVar)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "DupliObject":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*DupliObject, blk.Hdr.Count)
				for i := range bodies {
					body := new(DupliObject)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(DupliObject)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "DynamicPaintBrushSettings":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*DynamicPaintBrushSettings, blk.Hdr.Count)
				for i := range bodies {
					body := new(DynamicPaintBrushSettings)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(DynamicPaintBrushSettings)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "DynamicPaintCanvasSettings":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*DynamicPaintCanvasSettings, blk.Hdr.Count)
				for i := range bodies {
					body := new(DynamicPaintCanvasSettings)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(DynamicPaintCanvasSettings)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "DynamicPaintModifierData":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*DynamicPaintModifierData, blk.Hdr.Count)
				for i := range bodies {
					body := new(DynamicPaintModifierData)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(DynamicPaintModifierData)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "DynamicPaintSurface":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*DynamicPaintSurface, blk.Hdr.Count)
				for i := range bodies {
					body := new(DynamicPaintSurface)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(DynamicPaintSurface)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "EdgeSplitModifierData":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*EdgeSplitModifierData, blk.Hdr.Count)
				for i := range bodies {
					body := new(EdgeSplitModifierData)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(EdgeSplitModifierData)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "EditLatt":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*EditLatt, blk.Hdr.Count)
				for i := range bodies {
					body := new(EditLatt)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(EditLatt)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "EditNurb":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*EditNurb, blk.Hdr.Count)
				for i := range bodies {
					body := new(EditNurb)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(EditNurb)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "Editing":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*Editing, blk.Hdr.Count)
				for i := range bodies {
					body := new(Editing)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(Editing)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "Effect":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*Effect, blk.Hdr.Count)
				for i := range bodies {
					body := new(Effect)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(Effect)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "EffectorWeights":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*EffectorWeights, blk.Hdr.Count)
				for i := range bodies {
					body := new(EffectorWeights)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(EffectorWeights)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "EnvMap":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*EnvMap, blk.Hdr.Count)
				for i := range bodies {
					body := new(EnvMap)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(EnvMap)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "ExplodeModifierData":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*ExplodeModifierData, blk.Hdr.Count)
				for i := range bodies {
					body := new(ExplodeModifierData)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(ExplodeModifierData)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "FCM_EnvelopeData":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*FCM_EnvelopeData, blk.Hdr.Count)
				for i := range bodies {
					body := new(FCM_EnvelopeData)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(FCM_EnvelopeData)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "FCurve":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*FCurve, blk.Hdr.Count)
				for i := range bodies {
					body := new(FCurve)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(FCurve)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "FFMpegCodecData":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*FFMpegCodecData, blk.Hdr.Count)
				for i := range bodies {
					body := new(FFMpegCodecData)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(FFMpegCodecData)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "FMod_Cycles":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*FMod_Cycles, blk.Hdr.Count)
				for i := range bodies {
					body := new(FMod_Cycles)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(FMod_Cycles)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "FMod_Envelope":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*FMod_Envelope, blk.Hdr.Count)
				for i := range bodies {
					body := new(FMod_Envelope)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(FMod_Envelope)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "FMod_FunctionGenerator":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*FMod_FunctionGenerator, blk.Hdr.Count)
				for i := range bodies {
					body := new(FMod_FunctionGenerator)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(FMod_FunctionGenerator)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "FMod_Generator":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*FMod_Generator, blk.Hdr.Count)
				for i := range bodies {
					body := new(FMod_Generator)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(FMod_Generator)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "FMod_Limits":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*FMod_Limits, blk.Hdr.Count)
				for i := range bodies {
					body := new(FMod_Limits)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(FMod_Limits)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "FMod_Noise":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*FMod_Noise, blk.Hdr.Count)
				for i := range bodies {
					body := new(FMod_Noise)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(FMod_Noise)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "FMod_Python":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*FMod_Python, blk.Hdr.Count)
				for i := range bodies {
					body := new(FMod_Python)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(FMod_Python)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "FMod_Stepped":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*FMod_Stepped, blk.Hdr.Count)
				for i := range bodies {
					body := new(FMod_Stepped)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(FMod_Stepped)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "FModifier":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*FModifier, blk.Hdr.Count)
				for i := range bodies {
					body := new(FModifier)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(FModifier)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "FPoint":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*FPoint, blk.Hdr.Count)
				for i := range bodies {
					body := new(FPoint)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(FPoint)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "FileGlobal":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*FileGlobal, blk.Hdr.Count)
				for i := range bodies {
					body := new(FileGlobal)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(FileGlobal)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "FileSelectParams":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*FileSelectParams, blk.Hdr.Count)
				for i := range bodies {
					body := new(FileSelectParams)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(FileSelectParams)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "FluidVertexVelocity":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*FluidVertexVelocity, blk.Hdr.Count)
				for i := range bodies {
					body := new(FluidVertexVelocity)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(FluidVertexVelocity)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "FluidsimModifierData":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*FluidsimModifierData, blk.Hdr.Count)
				for i := range bodies {
					body := new(FluidsimModifierData)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(FluidsimModifierData)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "FluidsimSettings":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*FluidsimSettings, blk.Hdr.Count)
				for i := range bodies {
					body := new(FluidsimSettings)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(FluidsimSettings)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "FreestyleConfig":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*FreestyleConfig, blk.Hdr.Count)
				for i := range bodies {
					body := new(FreestyleConfig)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(FreestyleConfig)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "FreestyleEdge":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*FreestyleEdge, blk.Hdr.Count)
				for i := range bodies {
					body := new(FreestyleEdge)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(FreestyleEdge)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "FreestyleFace":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*FreestyleFace, blk.Hdr.Count)
				for i := range bodies {
					body := new(FreestyleFace)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(FreestyleFace)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "FreestyleLineSet":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*FreestyleLineSet, blk.Hdr.Count)
				for i := range bodies {
					body := new(FreestyleLineSet)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(FreestyleLineSet)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "FreestyleLineStyle":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*FreestyleLineStyle, blk.Hdr.Count)
				for i := range bodies {
					body := new(FreestyleLineStyle)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(FreestyleLineStyle)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "FreestyleModuleConfig":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*FreestyleModuleConfig, blk.Hdr.Count)
				for i := range bodies {
					body := new(FreestyleModuleConfig)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(FreestyleModuleConfig)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "GameData":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*GameData, blk.Hdr.Count)
				for i := range bodies {
					body := new(GameData)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(GameData)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "GameDome":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*GameDome, blk.Hdr.Count)
				for i := range bodies {
					body := new(GameDome)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(GameDome)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "GameFraming":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*GameFraming, blk.Hdr.Count)
				for i := range bodies {
					body := new(GameFraming)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(GameFraming)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "GameSettings":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*GameSettings, blk.Hdr.Count)
				for i := range bodies {
					body := new(GameSettings)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(GameSettings)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "GlowVars":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*GlowVars, blk.Hdr.Count)
				for i := range bodies {
					body := new(GlowVars)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(GlowVars)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "GridPaintMask":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*GridPaintMask, blk.Hdr.Count)
				for i := range bodies {
					body := new(GridPaintMask)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(GridPaintMask)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "Group":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*Group, blk.Hdr.Count)
				for i := range bodies {
					body := new(Group)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(Group)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "GroupObject":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*GroupObject, blk.Hdr.Count)
				for i := range bodies {
					body := new(GroupObject)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(GroupObject)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "HairKey":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*HairKey, blk.Hdr.Count)
				for i := range bodies {
					body := new(HairKey)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(HairKey)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "Histogram":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*Histogram, blk.Hdr.Count)
				for i := range bodies {
					body := new(Histogram)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(Histogram)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "HookModifierData":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*HookModifierData, blk.Hdr.Count)
				for i := range bodies {
					body := new(HookModifierData)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(HookModifierData)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "HueCorrectModifierData":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*HueCorrectModifierData, blk.Hdr.Count)
				for i := range bodies {
					body := new(HueCorrectModifierData)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(HueCorrectModifierData)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "ID":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*ID, blk.Hdr.Count)
				for i := range bodies {
					body := new(ID)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(ID)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "IDProperty":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*IDProperty, blk.Hdr.Count)
				for i := range bodies {
					body := new(IDProperty)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(IDProperty)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "IDPropertyData":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*IDPropertyData, blk.Hdr.Count)
				for i := range bodies {
					body := new(IDPropertyData)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(IDPropertyData)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "IdAdtTemplate":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*IdAdtTemplate, blk.Hdr.Count)
				for i := range bodies {
					body := new(IdAdtTemplate)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(IdAdtTemplate)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "Image":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*Image, blk.Hdr.Count)
				for i := range bodies {
					body := new(Image)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(Image)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "ImageFormatData":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*ImageFormatData, blk.Hdr.Count)
				for i := range bodies {
					body := new(ImageFormatData)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(ImageFormatData)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "ImagePaintSettings":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*ImagePaintSettings, blk.Hdr.Count)
				for i := range bodies {
					body := new(ImagePaintSettings)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(ImagePaintSettings)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "ImageUser":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*ImageUser, blk.Hdr.Count)
				for i := range bodies {
					body := new(ImageUser)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(ImageUser)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "Ipo":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*Ipo, blk.Hdr.Count)
				for i := range bodies {
					body := new(Ipo)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(Ipo)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "IpoCurve":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*IpoCurve, blk.Hdr.Count)
				for i := range bodies {
					body := new(IpoCurve)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(IpoCurve)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "IpoDriver":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*IpoDriver, blk.Hdr.Count)
				for i := range bodies {
					body := new(IpoDriver)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(IpoDriver)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "KS_Path":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*KS_Path, blk.Hdr.Count)
				for i := range bodies {
					body := new(KS_Path)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(KS_Path)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "Key":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*Key, blk.Hdr.Count)
				for i := range bodies {
					body := new(Key)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(Key)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "KeyBlock":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*KeyBlock, blk.Hdr.Count)
				for i := range bodies {
					body := new(KeyBlock)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(KeyBlock)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "KeyingSet":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*KeyingSet, blk.Hdr.Count)
				for i := range bodies {
					body := new(KeyingSet)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(KeyingSet)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "Lamp":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*Lamp, blk.Hdr.Count)
				for i := range bodies {
					body := new(Lamp)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(Lamp)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "LaplacianSmoothModifierData":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*LaplacianSmoothModifierData, blk.Hdr.Count)
				for i := range bodies {
					body := new(LaplacianSmoothModifierData)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(LaplacianSmoothModifierData)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "Lattice":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*Lattice, blk.Hdr.Count)
				for i := range bodies {
					body := new(Lattice)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(Lattice)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "LatticeModifierData":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*LatticeModifierData, blk.Hdr.Count)
				for i := range bodies {
					body := new(LatticeModifierData)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(LatticeModifierData)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "Library":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*Library, blk.Hdr.Count)
				for i := range bodies {
					body := new(Library)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(Library)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "LineStyleAlphaModifier_AlongStroke":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*LineStyleAlphaModifier_AlongStroke, blk.Hdr.Count)
				for i := range bodies {
					body := new(LineStyleAlphaModifier_AlongStroke)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(LineStyleAlphaModifier_AlongStroke)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "LineStyleAlphaModifier_DistanceFromCamera":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*LineStyleAlphaModifier_DistanceFromCamera, blk.Hdr.Count)
				for i := range bodies {
					body := new(LineStyleAlphaModifier_DistanceFromCamera)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(LineStyleAlphaModifier_DistanceFromCamera)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "LineStyleAlphaModifier_DistanceFromObject":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*LineStyleAlphaModifier_DistanceFromObject, blk.Hdr.Count)
				for i := range bodies {
					body := new(LineStyleAlphaModifier_DistanceFromObject)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(LineStyleAlphaModifier_DistanceFromObject)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "LineStyleAlphaModifier_Material":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*LineStyleAlphaModifier_Material, blk.Hdr.Count)
				for i := range bodies {
					body := new(LineStyleAlphaModifier_Material)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(LineStyleAlphaModifier_Material)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "LineStyleColorModifier_AlongStroke":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*LineStyleColorModifier_AlongStroke, blk.Hdr.Count)
				for i := range bodies {
					body := new(LineStyleColorModifier_AlongStroke)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(LineStyleColorModifier_AlongStroke)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "LineStyleColorModifier_DistanceFromCamera":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*LineStyleColorModifier_DistanceFromCamera, blk.Hdr.Count)
				for i := range bodies {
					body := new(LineStyleColorModifier_DistanceFromCamera)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(LineStyleColorModifier_DistanceFromCamera)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "LineStyleColorModifier_DistanceFromObject":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*LineStyleColorModifier_DistanceFromObject, blk.Hdr.Count)
				for i := range bodies {
					body := new(LineStyleColorModifier_DistanceFromObject)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(LineStyleColorModifier_DistanceFromObject)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "LineStyleColorModifier_Material":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*LineStyleColorModifier_Material, blk.Hdr.Count)
				for i := range bodies {
					body := new(LineStyleColorModifier_Material)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(LineStyleColorModifier_Material)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "LineStyleGeometryModifier_2DOffset":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*LineStyleGeometryModifier_2DOffset, blk.Hdr.Count)
				for i := range bodies {
					body := new(LineStyleGeometryModifier_2DOffset)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(LineStyleGeometryModifier_2DOffset)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "LineStyleGeometryModifier_2DTransform":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*LineStyleGeometryModifier_2DTransform, blk.Hdr.Count)
				for i := range bodies {
					body := new(LineStyleGeometryModifier_2DTransform)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(LineStyleGeometryModifier_2DTransform)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "LineStyleGeometryModifier_BackboneStretcher":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*LineStyleGeometryModifier_BackboneStretcher, blk.Hdr.Count)
				for i := range bodies {
					body := new(LineStyleGeometryModifier_BackboneStretcher)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(LineStyleGeometryModifier_BackboneStretcher)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "LineStyleGeometryModifier_BezierCurve":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*LineStyleGeometryModifier_BezierCurve, blk.Hdr.Count)
				for i := range bodies {
					body := new(LineStyleGeometryModifier_BezierCurve)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(LineStyleGeometryModifier_BezierCurve)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "LineStyleGeometryModifier_Blueprint":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*LineStyleGeometryModifier_Blueprint, blk.Hdr.Count)
				for i := range bodies {
					body := new(LineStyleGeometryModifier_Blueprint)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(LineStyleGeometryModifier_Blueprint)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "LineStyleGeometryModifier_GuidingLines":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*LineStyleGeometryModifier_GuidingLines, blk.Hdr.Count)
				for i := range bodies {
					body := new(LineStyleGeometryModifier_GuidingLines)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(LineStyleGeometryModifier_GuidingLines)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "LineStyleGeometryModifier_PerlinNoise1D":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*LineStyleGeometryModifier_PerlinNoise1D, blk.Hdr.Count)
				for i := range bodies {
					body := new(LineStyleGeometryModifier_PerlinNoise1D)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(LineStyleGeometryModifier_PerlinNoise1D)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "LineStyleGeometryModifier_PerlinNoise2D":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*LineStyleGeometryModifier_PerlinNoise2D, blk.Hdr.Count)
				for i := range bodies {
					body := new(LineStyleGeometryModifier_PerlinNoise2D)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(LineStyleGeometryModifier_PerlinNoise2D)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "LineStyleGeometryModifier_Polygonalization":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*LineStyleGeometryModifier_Polygonalization, blk.Hdr.Count)
				for i := range bodies {
					body := new(LineStyleGeometryModifier_Polygonalization)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(LineStyleGeometryModifier_Polygonalization)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "LineStyleGeometryModifier_Sampling":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*LineStyleGeometryModifier_Sampling, blk.Hdr.Count)
				for i := range bodies {
					body := new(LineStyleGeometryModifier_Sampling)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(LineStyleGeometryModifier_Sampling)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "LineStyleGeometryModifier_SinusDisplacement":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*LineStyleGeometryModifier_SinusDisplacement, blk.Hdr.Count)
				for i := range bodies {
					body := new(LineStyleGeometryModifier_SinusDisplacement)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(LineStyleGeometryModifier_SinusDisplacement)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "LineStyleGeometryModifier_SpatialNoise":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*LineStyleGeometryModifier_SpatialNoise, blk.Hdr.Count)
				for i := range bodies {
					body := new(LineStyleGeometryModifier_SpatialNoise)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(LineStyleGeometryModifier_SpatialNoise)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "LineStyleGeometryModifier_TipRemover":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*LineStyleGeometryModifier_TipRemover, blk.Hdr.Count)
				for i := range bodies {
					body := new(LineStyleGeometryModifier_TipRemover)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(LineStyleGeometryModifier_TipRemover)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "LineStyleModifier":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*LineStyleModifier, blk.Hdr.Count)
				for i := range bodies {
					body := new(LineStyleModifier)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(LineStyleModifier)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "LineStyleThicknessModifier_AlongStroke":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*LineStyleThicknessModifier_AlongStroke, blk.Hdr.Count)
				for i := range bodies {
					body := new(LineStyleThicknessModifier_AlongStroke)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(LineStyleThicknessModifier_AlongStroke)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "LineStyleThicknessModifier_Calligraphy":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*LineStyleThicknessModifier_Calligraphy, blk.Hdr.Count)
				for i := range bodies {
					body := new(LineStyleThicknessModifier_Calligraphy)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(LineStyleThicknessModifier_Calligraphy)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "LineStyleThicknessModifier_DistanceFromCamera":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*LineStyleThicknessModifier_DistanceFromCamera, blk.Hdr.Count)
				for i := range bodies {
					body := new(LineStyleThicknessModifier_DistanceFromCamera)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(LineStyleThicknessModifier_DistanceFromCamera)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "LineStyleThicknessModifier_DistanceFromObject":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*LineStyleThicknessModifier_DistanceFromObject, blk.Hdr.Count)
				for i := range bodies {
					body := new(LineStyleThicknessModifier_DistanceFromObject)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(LineStyleThicknessModifier_DistanceFromObject)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "LineStyleThicknessModifier_Material":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*LineStyleThicknessModifier_Material, blk.Hdr.Count)
				for i := range bodies {
					body := new(LineStyleThicknessModifier_Material)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(LineStyleThicknessModifier_Material)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "Link":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*Link, blk.Hdr.Count)
				for i := range bodies {
					body := new(Link)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(Link)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "LinkData":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*LinkData, blk.Hdr.Count)
				for i := range bodies {
					body := new(LinkData)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(LinkData)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "ListBase":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*ListBase, blk.Hdr.Count)
				for i := range bodies {
					body := new(ListBase)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(ListBase)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "MCol":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*MCol, blk.Hdr.Count)
				for i := range bodies {
					body := new(MCol)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(MCol)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "MDefCell":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*MDefCell, blk.Hdr.Count)
				for i := range bodies {
					body := new(MDefCell)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(MDefCell)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "MDefInfluence":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*MDefInfluence, blk.Hdr.Count)
				for i := range bodies {
					body := new(MDefInfluence)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(MDefInfluence)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "MDeformVert":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*MDeformVert, blk.Hdr.Count)
				for i := range bodies {
					body := new(MDeformVert)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(MDeformVert)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "MDeformWeight":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*MDeformWeight, blk.Hdr.Count)
				for i := range bodies {
					body := new(MDeformWeight)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(MDeformWeight)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "MDisps":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*MDisps, blk.Hdr.Count)
				for i := range bodies {
					body := new(MDisps)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(MDisps)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "MEdge":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*MEdge, blk.Hdr.Count)
				for i := range bodies {
					body := new(MEdge)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(MEdge)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "MFace":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*MFace, blk.Hdr.Count)
				for i := range bodies {
					body := new(MFace)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(MFace)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "MFloatProperty":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*MFloatProperty, blk.Hdr.Count)
				for i := range bodies {
					body := new(MFloatProperty)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(MFloatProperty)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "MIntProperty":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*MIntProperty, blk.Hdr.Count)
				for i := range bodies {
					body := new(MIntProperty)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(MIntProperty)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "MLoop":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*MLoop, blk.Hdr.Count)
				for i := range bodies {
					body := new(MLoop)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(MLoop)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "MLoopCol":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*MLoopCol, blk.Hdr.Count)
				for i := range bodies {
					body := new(MLoopCol)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(MLoopCol)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "MLoopUV":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*MLoopUV, blk.Hdr.Count)
				for i := range bodies {
					body := new(MLoopUV)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(MLoopUV)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "MPoly":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*MPoly, blk.Hdr.Count)
				for i := range bodies {
					body := new(MPoly)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(MPoly)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "MRecast":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*MRecast, blk.Hdr.Count)
				for i := range bodies {
					body := new(MRecast)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(MRecast)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "MSelect":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*MSelect, blk.Hdr.Count)
				for i := range bodies {
					body := new(MSelect)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(MSelect)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "MStringProperty":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*MStringProperty, blk.Hdr.Count)
				for i := range bodies {
					body := new(MStringProperty)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(MStringProperty)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "MTFace":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*MTFace, blk.Hdr.Count)
				for i := range bodies {
					body := new(MTFace)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(MTFace)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "MTex":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*MTex, blk.Hdr.Count)
				for i := range bodies {
					body := new(MTex)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(MTex)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "MTexPoly":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*MTexPoly, blk.Hdr.Count)
				for i := range bodies {
					body := new(MTexPoly)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(MTexPoly)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "MVert":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*MVert, blk.Hdr.Count)
				for i := range bodies {
					body := new(MVert)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(MVert)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "MVertSkin":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*MVertSkin, blk.Hdr.Count)
				for i := range bodies {
					body := new(MVertSkin)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(MVertSkin)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "MappingInfoModifierData":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*MappingInfoModifierData, blk.Hdr.Count)
				for i := range bodies {
					body := new(MappingInfoModifierData)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(MappingInfoModifierData)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "Mask":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*Mask, blk.Hdr.Count)
				for i := range bodies {
					body := new(Mask)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(Mask)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "MaskLayer":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*MaskLayer, blk.Hdr.Count)
				for i := range bodies {
					body := new(MaskLayer)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(MaskLayer)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "MaskLayerShape":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*MaskLayerShape, blk.Hdr.Count)
				for i := range bodies {
					body := new(MaskLayerShape)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(MaskLayerShape)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "MaskModifierData":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*MaskModifierData, blk.Hdr.Count)
				for i := range bodies {
					body := new(MaskModifierData)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(MaskModifierData)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "MaskParent":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*MaskParent, blk.Hdr.Count)
				for i := range bodies {
					body := new(MaskParent)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(MaskParent)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "MaskSpaceInfo":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*MaskSpaceInfo, blk.Hdr.Count)
				for i := range bodies {
					body := new(MaskSpaceInfo)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(MaskSpaceInfo)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "MaskSpline":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*MaskSpline, blk.Hdr.Count)
				for i := range bodies {
					body := new(MaskSpline)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(MaskSpline)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "MaskSplinePoint":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*MaskSplinePoint, blk.Hdr.Count)
				for i := range bodies {
					body := new(MaskSplinePoint)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(MaskSplinePoint)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "MaskSplinePointUW":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*MaskSplinePointUW, blk.Hdr.Count)
				for i := range bodies {
					body := new(MaskSplinePointUW)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(MaskSplinePointUW)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "Material":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*Material, blk.Hdr.Count)
				for i := range bodies {
					body := new(Material)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(Material)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "Mesh":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*Mesh, blk.Hdr.Count)
				for i := range bodies {
					body := new(Mesh)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(Mesh)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "MeshCacheModifierData":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*MeshCacheModifierData, blk.Hdr.Count)
				for i := range bodies {
					body := new(MeshCacheModifierData)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(MeshCacheModifierData)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "MeshDeformModifierData":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*MeshDeformModifierData, blk.Hdr.Count)
				for i := range bodies {
					body := new(MeshDeformModifierData)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(MeshDeformModifierData)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "MeshStatVis":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*MeshStatVis, blk.Hdr.Count)
				for i := range bodies {
					body := new(MeshStatVis)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(MeshStatVis)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "MetaBall":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*MetaBall, blk.Hdr.Count)
				for i := range bodies {
					body := new(MetaBall)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(MetaBall)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "MetaElem":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*MetaElem, blk.Hdr.Count)
				for i := range bodies {
					body := new(MetaElem)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(MetaElem)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "MetaStack":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*MetaStack, blk.Hdr.Count)
				for i := range bodies {
					body := new(MetaStack)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(MetaStack)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "MirrorModifierData":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*MirrorModifierData, blk.Hdr.Count)
				for i := range bodies {
					body := new(MirrorModifierData)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(MirrorModifierData)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "ModifierData":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*ModifierData, blk.Hdr.Count)
				for i := range bodies {
					body := new(ModifierData)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(ModifierData)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "MovieClip":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*MovieClip, blk.Hdr.Count)
				for i := range bodies {
					body := new(MovieClip)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(MovieClip)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "MovieClipProxy":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*MovieClipProxy, blk.Hdr.Count)
				for i := range bodies {
					body := new(MovieClipProxy)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(MovieClipProxy)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "MovieClipScopes":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*MovieClipScopes, blk.Hdr.Count)
				for i := range bodies {
					body := new(MovieClipScopes)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(MovieClipScopes)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "MovieClipUser":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*MovieClipUser, blk.Hdr.Count)
				for i := range bodies {
					body := new(MovieClipUser)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(MovieClipUser)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "MovieReconstructedCamera":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*MovieReconstructedCamera, blk.Hdr.Count)
				for i := range bodies {
					body := new(MovieReconstructedCamera)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(MovieReconstructedCamera)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "MovieTracking":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*MovieTracking, blk.Hdr.Count)
				for i := range bodies {
					body := new(MovieTracking)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(MovieTracking)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "MovieTrackingCamera":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*MovieTrackingCamera, blk.Hdr.Count)
				for i := range bodies {
					body := new(MovieTrackingCamera)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(MovieTrackingCamera)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "MovieTrackingDopesheet":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*MovieTrackingDopesheet, blk.Hdr.Count)
				for i := range bodies {
					body := new(MovieTrackingDopesheet)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(MovieTrackingDopesheet)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "MovieTrackingDopesheetChannel":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*MovieTrackingDopesheetChannel, blk.Hdr.Count)
				for i := range bodies {
					body := new(MovieTrackingDopesheetChannel)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(MovieTrackingDopesheetChannel)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "MovieTrackingDopesheetCoverageSegment":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*MovieTrackingDopesheetCoverageSegment, blk.Hdr.Count)
				for i := range bodies {
					body := new(MovieTrackingDopesheetCoverageSegment)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(MovieTrackingDopesheetCoverageSegment)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "MovieTrackingMarker":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*MovieTrackingMarker, blk.Hdr.Count)
				for i := range bodies {
					body := new(MovieTrackingMarker)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(MovieTrackingMarker)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "MovieTrackingObject":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*MovieTrackingObject, blk.Hdr.Count)
				for i := range bodies {
					body := new(MovieTrackingObject)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(MovieTrackingObject)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "MovieTrackingReconstruction":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*MovieTrackingReconstruction, blk.Hdr.Count)
				for i := range bodies {
					body := new(MovieTrackingReconstruction)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(MovieTrackingReconstruction)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "MovieTrackingSettings":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*MovieTrackingSettings, blk.Hdr.Count)
				for i := range bodies {
					body := new(MovieTrackingSettings)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(MovieTrackingSettings)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "MovieTrackingStabilization":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*MovieTrackingStabilization, blk.Hdr.Count)
				for i := range bodies {
					body := new(MovieTrackingStabilization)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(MovieTrackingStabilization)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "MovieTrackingStats":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*MovieTrackingStats, blk.Hdr.Count)
				for i := range bodies {
					body := new(MovieTrackingStats)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(MovieTrackingStats)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "MovieTrackingTrack":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*MovieTrackingTrack, blk.Hdr.Count)
				for i := range bodies {
					body := new(MovieTrackingTrack)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(MovieTrackingTrack)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "Multires":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*Multires, blk.Hdr.Count)
				for i := range bodies {
					body := new(Multires)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(Multires)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "MultiresCol":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*MultiresCol, blk.Hdr.Count)
				for i := range bodies {
					body := new(MultiresCol)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(MultiresCol)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "MultiresColFace":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*MultiresColFace, blk.Hdr.Count)
				for i := range bodies {
					body := new(MultiresColFace)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(MultiresColFace)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "MultiresEdge":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*MultiresEdge, blk.Hdr.Count)
				for i := range bodies {
					body := new(MultiresEdge)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(MultiresEdge)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "MultiresFace":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*MultiresFace, blk.Hdr.Count)
				for i := range bodies {
					body := new(MultiresFace)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(MultiresFace)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "MultiresLevel":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*MultiresLevel, blk.Hdr.Count)
				for i := range bodies {
					body := new(MultiresLevel)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(MultiresLevel)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "MultiresModifierData":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*MultiresModifierData, blk.Hdr.Count)
				for i := range bodies {
					body := new(MultiresModifierData)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(MultiresModifierData)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "NlaStrip":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*NlaStrip, blk.Hdr.Count)
				for i := range bodies {
					body := new(NlaStrip)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(NlaStrip)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "NlaTrack":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*NlaTrack, blk.Hdr.Count)
				for i := range bodies {
					body := new(NlaTrack)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(NlaTrack)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "NodeBilateralBlurData":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*NodeBilateralBlurData, blk.Hdr.Count)
				for i := range bodies {
					body := new(NodeBilateralBlurData)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(NodeBilateralBlurData)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "NodeBlurData":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*NodeBlurData, blk.Hdr.Count)
				for i := range bodies {
					body := new(NodeBlurData)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(NodeBlurData)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "NodeBokehImage":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*NodeBokehImage, blk.Hdr.Count)
				for i := range bodies {
					body := new(NodeBokehImage)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(NodeBokehImage)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "NodeBoxMask":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*NodeBoxMask, blk.Hdr.Count)
				for i := range bodies {
					body := new(NodeBoxMask)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(NodeBoxMask)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "NodeChroma":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*NodeChroma, blk.Hdr.Count)
				for i := range bodies {
					body := new(NodeChroma)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(NodeChroma)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "NodeColorBalance":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*NodeColorBalance, blk.Hdr.Count)
				for i := range bodies {
					body := new(NodeColorBalance)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(NodeColorBalance)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "NodeColorCorrection":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*NodeColorCorrection, blk.Hdr.Count)
				for i := range bodies {
					body := new(NodeColorCorrection)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(NodeColorCorrection)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "NodeColorspill":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*NodeColorspill, blk.Hdr.Count)
				for i := range bodies {
					body := new(NodeColorspill)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(NodeColorspill)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "NodeDBlurData":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*NodeDBlurData, blk.Hdr.Count)
				for i := range bodies {
					body := new(NodeDBlurData)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(NodeDBlurData)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "NodeDefocus":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*NodeDefocus, blk.Hdr.Count)
				for i := range bodies {
					body := new(NodeDefocus)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(NodeDefocus)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "NodeDilateErode":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*NodeDilateErode, blk.Hdr.Count)
				for i := range bodies {
					body := new(NodeDilateErode)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(NodeDilateErode)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "NodeEllipseMask":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*NodeEllipseMask, blk.Hdr.Count)
				for i := range bodies {
					body := new(NodeEllipseMask)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(NodeEllipseMask)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "NodeFrame":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*NodeFrame, blk.Hdr.Count)
				for i := range bodies {
					body := new(NodeFrame)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(NodeFrame)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "NodeGeometry":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*NodeGeometry, blk.Hdr.Count)
				for i := range bodies {
					body := new(NodeGeometry)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(NodeGeometry)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "NodeGlare":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*NodeGlare, blk.Hdr.Count)
				for i := range bodies {
					body := new(NodeGlare)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(NodeGlare)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "NodeHueSat":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*NodeHueSat, blk.Hdr.Count)
				for i := range bodies {
					body := new(NodeHueSat)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(NodeHueSat)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "NodeImageAnim":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*NodeImageAnim, blk.Hdr.Count)
				for i := range bodies {
					body := new(NodeImageAnim)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(NodeImageAnim)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "NodeImageFile":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*NodeImageFile, blk.Hdr.Count)
				for i := range bodies {
					body := new(NodeImageFile)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(NodeImageFile)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "NodeImageLayer":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*NodeImageLayer, blk.Hdr.Count)
				for i := range bodies {
					body := new(NodeImageLayer)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(NodeImageLayer)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "NodeImageMultiFile":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*NodeImageMultiFile, blk.Hdr.Count)
				for i := range bodies {
					body := new(NodeImageMultiFile)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(NodeImageMultiFile)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "NodeImageMultiFileSocket":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*NodeImageMultiFileSocket, blk.Hdr.Count)
				for i := range bodies {
					body := new(NodeImageMultiFileSocket)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(NodeImageMultiFileSocket)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "NodeKeyingData":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*NodeKeyingData, blk.Hdr.Count)
				for i := range bodies {
					body := new(NodeKeyingData)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(NodeKeyingData)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "NodeKeyingScreenData":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*NodeKeyingScreenData, blk.Hdr.Count)
				for i := range bodies {
					body := new(NodeKeyingScreenData)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(NodeKeyingScreenData)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "NodeLensDist":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*NodeLensDist, blk.Hdr.Count)
				for i := range bodies {
					body := new(NodeLensDist)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(NodeLensDist)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "NodeMask":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*NodeMask, blk.Hdr.Count)
				for i := range bodies {
					body := new(NodeMask)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(NodeMask)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "NodeScriptDict":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*NodeScriptDict, blk.Hdr.Count)
				for i := range bodies {
					body := new(NodeScriptDict)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(NodeScriptDict)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "NodeShaderAttribute":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*NodeShaderAttribute, blk.Hdr.Count)
				for i := range bodies {
					body := new(NodeShaderAttribute)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(NodeShaderAttribute)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "NodeShaderNormalMap":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*NodeShaderNormalMap, blk.Hdr.Count)
				for i := range bodies {
					body := new(NodeShaderNormalMap)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(NodeShaderNormalMap)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "NodeShaderScript":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*NodeShaderScript, blk.Hdr.Count)
				for i := range bodies {
					body := new(NodeShaderScript)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(NodeShaderScript)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "NodeShaderTangent":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*NodeShaderTangent, blk.Hdr.Count)
				for i := range bodies {
					body := new(NodeShaderTangent)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(NodeShaderTangent)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "NodeTexBase":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*NodeTexBase, blk.Hdr.Count)
				for i := range bodies {
					body := new(NodeTexBase)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(NodeTexBase)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "NodeTexBrick":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*NodeTexBrick, blk.Hdr.Count)
				for i := range bodies {
					body := new(NodeTexBrick)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(NodeTexBrick)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "NodeTexChecker":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*NodeTexChecker, blk.Hdr.Count)
				for i := range bodies {
					body := new(NodeTexChecker)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(NodeTexChecker)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "NodeTexEnvironment":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*NodeTexEnvironment, blk.Hdr.Count)
				for i := range bodies {
					body := new(NodeTexEnvironment)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(NodeTexEnvironment)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "NodeTexGradient":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*NodeTexGradient, blk.Hdr.Count)
				for i := range bodies {
					body := new(NodeTexGradient)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(NodeTexGradient)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "NodeTexImage":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*NodeTexImage, blk.Hdr.Count)
				for i := range bodies {
					body := new(NodeTexImage)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(NodeTexImage)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "NodeTexMagic":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*NodeTexMagic, blk.Hdr.Count)
				for i := range bodies {
					body := new(NodeTexMagic)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(NodeTexMagic)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "NodeTexMusgrave":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*NodeTexMusgrave, blk.Hdr.Count)
				for i := range bodies {
					body := new(NodeTexMusgrave)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(NodeTexMusgrave)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "NodeTexNoise":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*NodeTexNoise, blk.Hdr.Count)
				for i := range bodies {
					body := new(NodeTexNoise)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(NodeTexNoise)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "NodeTexSky":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*NodeTexSky, blk.Hdr.Count)
				for i := range bodies {
					body := new(NodeTexSky)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(NodeTexSky)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "NodeTexVoronoi":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*NodeTexVoronoi, blk.Hdr.Count)
				for i := range bodies {
					body := new(NodeTexVoronoi)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(NodeTexVoronoi)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "NodeTexWave":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*NodeTexWave, blk.Hdr.Count)
				for i := range bodies {
					body := new(NodeTexWave)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(NodeTexWave)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "NodeTonemap":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*NodeTonemap, blk.Hdr.Count)
				for i := range bodies {
					body := new(NodeTonemap)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(NodeTonemap)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "NodeTrackPosData":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*NodeTrackPosData, blk.Hdr.Count)
				for i := range bodies {
					body := new(NodeTrackPosData)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(NodeTrackPosData)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "NodeTranslateData":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*NodeTranslateData, blk.Hdr.Count)
				for i := range bodies {
					body := new(NodeTranslateData)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(NodeTranslateData)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "NodeTwoFloats":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*NodeTwoFloats, blk.Hdr.Count)
				for i := range bodies {
					body := new(NodeTwoFloats)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(NodeTwoFloats)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "NodeTwoXYs":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*NodeTwoXYs, blk.Hdr.Count)
				for i := range bodies {
					body := new(NodeTwoXYs)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(NodeTwoXYs)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "NodeVertexCol":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*NodeVertexCol, blk.Hdr.Count)
				for i := range bodies {
					body := new(NodeVertexCol)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(NodeVertexCol)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "Nurb":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*Nurb, blk.Hdr.Count)
				for i := range bodies {
					body := new(Nurb)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(Nurb)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "ObHook":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*ObHook, blk.Hdr.Count)
				for i := range bodies {
					body := new(ObHook)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(ObHook)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "Object":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*Object, blk.Hdr.Count)
				for i := range bodies {
					body := new(Object)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(Object)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "OceanModifierData":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*OceanModifierData, blk.Hdr.Count)
				for i := range bodies {
					body := new(OceanModifierData)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(OceanModifierData)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "OceanTex":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*OceanTex, blk.Hdr.Count)
				for i := range bodies {
					body := new(OceanTex)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(OceanTex)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "OrigSpaceFace":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*OrigSpaceFace, blk.Hdr.Count)
				for i := range bodies {
					body := new(OrigSpaceFace)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(OrigSpaceFace)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "OrigSpaceLoop":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*OrigSpaceLoop, blk.Hdr.Count)
				for i := range bodies {
					body := new(OrigSpaceLoop)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(OrigSpaceLoop)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "PTCacheExtra":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*PTCacheExtra, blk.Hdr.Count)
				for i := range bodies {
					body := new(PTCacheExtra)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(PTCacheExtra)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "PTCacheMem":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*PTCacheMem, blk.Hdr.Count)
				for i := range bodies {
					body := new(PTCacheMem)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(PTCacheMem)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "PackedFile":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*PackedFile, blk.Hdr.Count)
				for i := range bodies {
					body := new(PackedFile)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(PackedFile)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "Paint":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*Paint, blk.Hdr.Count)
				for i := range bodies {
					body := new(Paint)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(Paint)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "Panel":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*Panel, blk.Hdr.Count)
				for i := range bodies {
					body := new(Panel)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(Panel)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "PartDeflect":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*PartDeflect, blk.Hdr.Count)
				for i := range bodies {
					body := new(PartDeflect)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(PartDeflect)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "PartEff":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*PartEff, blk.Hdr.Count)
				for i := range bodies {
					body := new(PartEff)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(PartEff)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "ParticleBrushData":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*ParticleBrushData, blk.Hdr.Count)
				for i := range bodies {
					body := new(ParticleBrushData)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(ParticleBrushData)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "ParticleData":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*ParticleData, blk.Hdr.Count)
				for i := range bodies {
					body := new(ParticleData)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(ParticleData)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "ParticleDupliWeight":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*ParticleDupliWeight, blk.Hdr.Count)
				for i := range bodies {
					body := new(ParticleDupliWeight)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(ParticleDupliWeight)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "ParticleEditSettings":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*ParticleEditSettings, blk.Hdr.Count)
				for i := range bodies {
					body := new(ParticleEditSettings)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(ParticleEditSettings)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "ParticleInstanceModifierData":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*ParticleInstanceModifierData, blk.Hdr.Count)
				for i := range bodies {
					body := new(ParticleInstanceModifierData)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(ParticleInstanceModifierData)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "ParticleKey":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*ParticleKey, blk.Hdr.Count)
				for i := range bodies {
					body := new(ParticleKey)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(ParticleKey)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "ParticleSettings":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*ParticleSettings, blk.Hdr.Count)
				for i := range bodies {
					body := new(ParticleSettings)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(ParticleSettings)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "ParticleSpring":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*ParticleSpring, blk.Hdr.Count)
				for i := range bodies {
					body := new(ParticleSpring)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(ParticleSpring)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "ParticleSystem":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*ParticleSystem, blk.Hdr.Count)
				for i := range bodies {
					body := new(ParticleSystem)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(ParticleSystem)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "ParticleSystemModifierData":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*ParticleSystemModifierData, blk.Hdr.Count)
				for i := range bodies {
					body := new(ParticleSystemModifierData)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(ParticleSystemModifierData)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "ParticleTarget":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*ParticleTarget, blk.Hdr.Count)
				for i := range bodies {
					body := new(ParticleTarget)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(ParticleTarget)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "PhysicsSettings":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*PhysicsSettings, blk.Hdr.Count)
				for i := range bodies {
					body := new(PhysicsSettings)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(PhysicsSettings)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "PointCache":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*PointCache, blk.Hdr.Count)
				for i := range bodies {
					body := new(PointCache)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(PointCache)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "PointDensity":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*PointDensity, blk.Hdr.Count)
				for i := range bodies {
					body := new(PointDensity)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(PointDensity)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "PreviewImage":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*PreviewImage, blk.Hdr.Count)
				for i := range bodies {
					body := new(PreviewImage)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(PreviewImage)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "QuicktimeCodecData":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*QuicktimeCodecData, blk.Hdr.Count)
				for i := range bodies {
					body := new(QuicktimeCodecData)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(QuicktimeCodecData)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "QuicktimeCodecSettings":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*QuicktimeCodecSettings, blk.Hdr.Count)
				for i := range bodies {
					body := new(QuicktimeCodecSettings)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(QuicktimeCodecSettings)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "RecastData":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*RecastData, blk.Hdr.Count)
				for i := range bodies {
					body := new(RecastData)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(RecastData)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "RegionView3D":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*RegionView3D, blk.Hdr.Count)
				for i := range bodies {
					body := new(RegionView3D)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(RegionView3D)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "RemeshModifierData":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*RemeshModifierData, blk.Hdr.Count)
				for i := range bodies {
					body := new(RemeshModifierData)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(RemeshModifierData)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "RenderData":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*RenderData, blk.Hdr.Count)
				for i := range bodies {
					body := new(RenderData)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(RenderData)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "RenderProfile":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*RenderProfile, blk.Hdr.Count)
				for i := range bodies {
					body := new(RenderProfile)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(RenderProfile)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "ReportList":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*ReportList, blk.Hdr.Count)
				for i := range bodies {
					body := new(ReportList)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(ReportList)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "RigidBodyCon":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*RigidBodyCon, blk.Hdr.Count)
				for i := range bodies {
					body := new(RigidBodyCon)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(RigidBodyCon)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "RigidBodyOb":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*RigidBodyOb, blk.Hdr.Count)
				for i := range bodies {
					body := new(RigidBodyOb)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(RigidBodyOb)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "RigidBodyWorld":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*RigidBodyWorld, blk.Hdr.Count)
				for i := range bodies {
					body := new(RigidBodyWorld)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(RigidBodyWorld)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "SBVertex":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*SBVertex, blk.Hdr.Count)
				for i := range bodies {
					body := new(SBVertex)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(SBVertex)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "SPHFluidSettings":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*SPHFluidSettings, blk.Hdr.Count)
				for i := range bodies {
					body := new(SPHFluidSettings)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(SPHFluidSettings)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "Scene":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*Scene, blk.Hdr.Count)
				for i := range bodies {
					body := new(Scene)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(Scene)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "SceneRenderLayer":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*SceneRenderLayer, blk.Hdr.Count)
				for i := range bodies {
					body := new(SceneRenderLayer)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(SceneRenderLayer)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "Scopes":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*Scopes, blk.Hdr.Count)
				for i := range bodies {
					body := new(Scopes)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(Scopes)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "ScrArea":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*ScrArea, blk.Hdr.Count)
				for i := range bodies {
					body := new(ScrArea)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(ScrArea)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "ScrEdge":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*ScrEdge, blk.Hdr.Count)
				for i := range bodies {
					body := new(ScrEdge)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(ScrEdge)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "ScrVert":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*ScrVert, blk.Hdr.Count)
				for i := range bodies {
					body := new(ScrVert)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(ScrVert)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "ScrewModifierData":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*ScrewModifierData, blk.Hdr.Count)
				for i := range bodies {
					body := new(ScrewModifierData)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(ScrewModifierData)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "Script":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*Script, blk.Hdr.Count)
				for i := range bodies {
					body := new(Script)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(Script)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "Sculpt":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*Sculpt, blk.Hdr.Count)
				for i := range bodies {
					body := new(Sculpt)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(Sculpt)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "Sequence":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*Sequence, blk.Hdr.Count)
				for i := range bodies {
					body := new(Sequence)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(Sequence)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "SequenceModifierData":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*SequenceModifierData, blk.Hdr.Count)
				for i := range bodies {
					body := new(SequenceModifierData)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(SequenceModifierData)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "SequencerScopes":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*SequencerScopes, blk.Hdr.Count)
				for i := range bodies {
					body := new(SequencerScopes)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(SequencerScopes)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "ShapeKeyModifierData":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*ShapeKeyModifierData, blk.Hdr.Count)
				for i := range bodies {
					body := new(ShapeKeyModifierData)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(ShapeKeyModifierData)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "ShrinkwrapModifierData":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*ShrinkwrapModifierData, blk.Hdr.Count)
				for i := range bodies {
					body := new(ShrinkwrapModifierData)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(ShrinkwrapModifierData)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "SimpleDeformModifierData":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*SimpleDeformModifierData, blk.Hdr.Count)
				for i := range bodies {
					body := new(SimpleDeformModifierData)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(SimpleDeformModifierData)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "SkinModifierData":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*SkinModifierData, blk.Hdr.Count)
				for i := range bodies {
					body := new(SkinModifierData)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(SkinModifierData)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "SmokeCollSettings":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*SmokeCollSettings, blk.Hdr.Count)
				for i := range bodies {
					body := new(SmokeCollSettings)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(SmokeCollSettings)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "SmokeDomainSettings":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*SmokeDomainSettings, blk.Hdr.Count)
				for i := range bodies {
					body := new(SmokeDomainSettings)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(SmokeDomainSettings)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "SmokeFlowSettings":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*SmokeFlowSettings, blk.Hdr.Count)
				for i := range bodies {
					body := new(SmokeFlowSettings)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(SmokeFlowSettings)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "SmokeModifierData":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*SmokeModifierData, blk.Hdr.Count)
				for i := range bodies {
					body := new(SmokeModifierData)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(SmokeModifierData)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "SmoothModifierData":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*SmoothModifierData, blk.Hdr.Count)
				for i := range bodies {
					body := new(SmoothModifierData)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(SmoothModifierData)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "SoftBody":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*SoftBody, blk.Hdr.Count)
				for i := range bodies {
					body := new(SoftBody)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(SoftBody)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "SoftbodyModifierData":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*SoftbodyModifierData, blk.Hdr.Count)
				for i := range bodies {
					body := new(SoftbodyModifierData)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(SoftbodyModifierData)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "SolidColorVars":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*SolidColorVars, blk.Hdr.Count)
				for i := range bodies {
					body := new(SolidColorVars)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(SolidColorVars)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "SolidLight":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*SolidLight, blk.Hdr.Count)
				for i := range bodies {
					body := new(SolidLight)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(SolidLight)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "SolidifyModifierData":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*SolidifyModifierData, blk.Hdr.Count)
				for i := range bodies {
					body := new(SolidifyModifierData)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(SolidifyModifierData)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "Sound3D":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*Sound3D, blk.Hdr.Count)
				for i := range bodies {
					body := new(Sound3D)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(Sound3D)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "SpaceAction":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*SpaceAction, blk.Hdr.Count)
				for i := range bodies {
					body := new(SpaceAction)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(SpaceAction)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "SpaceButs":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*SpaceButs, blk.Hdr.Count)
				for i := range bodies {
					body := new(SpaceButs)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(SpaceButs)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "SpaceClip":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*SpaceClip, blk.Hdr.Count)
				for i := range bodies {
					body := new(SpaceClip)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(SpaceClip)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "SpaceConsole":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*SpaceConsole, blk.Hdr.Count)
				for i := range bodies {
					body := new(SpaceConsole)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(SpaceConsole)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "SpaceFile":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*SpaceFile, blk.Hdr.Count)
				for i := range bodies {
					body := new(SpaceFile)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(SpaceFile)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "SpaceImage":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*SpaceImage, blk.Hdr.Count)
				for i := range bodies {
					body := new(SpaceImage)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(SpaceImage)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "SpaceInfo":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*SpaceInfo, blk.Hdr.Count)
				for i := range bodies {
					body := new(SpaceInfo)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(SpaceInfo)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "SpaceIpo":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*SpaceIpo, blk.Hdr.Count)
				for i := range bodies {
					body := new(SpaceIpo)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(SpaceIpo)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "SpaceLink":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*SpaceLink, blk.Hdr.Count)
				for i := range bodies {
					body := new(SpaceLink)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(SpaceLink)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "SpaceLogic":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*SpaceLogic, blk.Hdr.Count)
				for i := range bodies {
					body := new(SpaceLogic)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(SpaceLogic)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "SpaceNla":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*SpaceNla, blk.Hdr.Count)
				for i := range bodies {
					body := new(SpaceNla)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(SpaceNla)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "SpaceNode":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*SpaceNode, blk.Hdr.Count)
				for i := range bodies {
					body := new(SpaceNode)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(SpaceNode)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "SpaceOops":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*SpaceOops, blk.Hdr.Count)
				for i := range bodies {
					body := new(SpaceOops)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(SpaceOops)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "SpaceScript":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*SpaceScript, blk.Hdr.Count)
				for i := range bodies {
					body := new(SpaceScript)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(SpaceScript)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "SpaceSeq":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*SpaceSeq, blk.Hdr.Count)
				for i := range bodies {
					body := new(SpaceSeq)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(SpaceSeq)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "SpaceText":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*SpaceText, blk.Hdr.Count)
				for i := range bodies {
					body := new(SpaceText)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(SpaceText)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "SpaceTime":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*SpaceTime, blk.Hdr.Count)
				for i := range bodies {
					body := new(SpaceTime)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(SpaceTime)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "SpaceTimeCache":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*SpaceTimeCache, blk.Hdr.Count)
				for i := range bodies {
					body := new(SpaceTimeCache)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(SpaceTimeCache)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "SpaceUserPref":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*SpaceUserPref, blk.Hdr.Count)
				for i := range bodies {
					body := new(SpaceUserPref)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(SpaceUserPref)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "Speaker":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*Speaker, blk.Hdr.Count)
				for i := range bodies {
					body := new(Speaker)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(Speaker)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "SpeedControlVars":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*SpeedControlVars, blk.Hdr.Count)
				for i := range bodies {
					body := new(SpeedControlVars)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(SpeedControlVars)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "Strip":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*Strip, blk.Hdr.Count)
				for i := range bodies {
					body := new(Strip)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(Strip)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "StripColorBalance":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*StripColorBalance, blk.Hdr.Count)
				for i := range bodies {
					body := new(StripColorBalance)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(StripColorBalance)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "StripCrop":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*StripCrop, blk.Hdr.Count)
				for i := range bodies {
					body := new(StripCrop)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(StripCrop)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "StripElem":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*StripElem, blk.Hdr.Count)
				for i := range bodies {
					body := new(StripElem)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(StripElem)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "StripProxy":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*StripProxy, blk.Hdr.Count)
				for i := range bodies {
					body := new(StripProxy)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(StripProxy)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "StripTransform":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*StripTransform, blk.Hdr.Count)
				for i := range bodies {
					body := new(StripTransform)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(StripTransform)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "SubsurfModifierData":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*SubsurfModifierData, blk.Hdr.Count)
				for i := range bodies {
					body := new(SubsurfModifierData)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(SubsurfModifierData)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "SurfaceModifierData":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*SurfaceModifierData, blk.Hdr.Count)
				for i := range bodies {
					body := new(SurfaceModifierData)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(SurfaceModifierData)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "TFace":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*TFace, blk.Hdr.Count)
				for i := range bodies {
					body := new(TFace)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(TFace)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "Tex":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*Tex, blk.Hdr.Count)
				for i := range bodies {
					body := new(Tex)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(Tex)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "TexMapping":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*TexMapping, blk.Hdr.Count)
				for i := range bodies {
					body := new(TexMapping)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(TexMapping)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "TexNodeOutput":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*TexNodeOutput, blk.Hdr.Count)
				for i := range bodies {
					body := new(TexNodeOutput)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(TexNodeOutput)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "Text":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*Text, blk.Hdr.Count)
				for i := range bodies {
					body := new(Text)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(Text)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "TextBox":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*TextBox, blk.Hdr.Count)
				for i := range bodies {
					body := new(TextBox)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(TextBox)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "TextLine":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*TextLine, blk.Hdr.Count)
				for i := range bodies {
					body := new(TextLine)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(TextLine)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "ThemeSpace":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*ThemeSpace, blk.Hdr.Count)
				for i := range bodies {
					body := new(ThemeSpace)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(ThemeSpace)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "ThemeUI":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*ThemeUI, blk.Hdr.Count)
				for i := range bodies {
					body := new(ThemeUI)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(ThemeUI)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "ThemeWireColor":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*ThemeWireColor, blk.Hdr.Count)
				for i := range bodies {
					body := new(ThemeWireColor)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(ThemeWireColor)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "TimeMarker":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*TimeMarker, blk.Hdr.Count)
				for i := range bodies {
					body := new(TimeMarker)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(TimeMarker)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "ToolSettings":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*ToolSettings, blk.Hdr.Count)
				for i := range bodies {
					body := new(ToolSettings)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(ToolSettings)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "TransformOrientation":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*TransformOrientation, blk.Hdr.Count)
				for i := range bodies {
					body := new(TransformOrientation)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(TransformOrientation)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "TransformVars":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*TransformVars, blk.Hdr.Count)
				for i := range bodies {
					body := new(TransformVars)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(TransformVars)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "TreeStore":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*TreeStore, blk.Hdr.Count)
				for i := range bodies {
					body := new(TreeStore)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(TreeStore)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "TreeStoreElem":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*TreeStoreElem, blk.Hdr.Count)
				for i := range bodies {
					body := new(TreeStoreElem)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(TreeStoreElem)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "TriangulateModifierData":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*TriangulateModifierData, blk.Hdr.Count)
				for i := range bodies {
					body := new(TriangulateModifierData)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(TriangulateModifierData)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "UVProjectModifierData":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*UVProjectModifierData, blk.Hdr.Count)
				for i := range bodies {
					body := new(UVProjectModifierData)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(UVProjectModifierData)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "UVWarpModifierData":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*UVWarpModifierData, blk.Hdr.Count)
				for i := range bodies {
					body := new(UVWarpModifierData)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(UVWarpModifierData)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "UnifiedPaintSettings":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*UnifiedPaintSettings, blk.Hdr.Count)
				for i := range bodies {
					body := new(UnifiedPaintSettings)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(UnifiedPaintSettings)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "UnitSettings":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*UnitSettings, blk.Hdr.Count)
				for i := range bodies {
					body := new(UnitSettings)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(UnitSettings)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "UserDef":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*UserDef, blk.Hdr.Count)
				for i := range bodies {
					body := new(UserDef)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(UserDef)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "UvSculpt":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*UvSculpt, blk.Hdr.Count)
				for i := range bodies {
					body := new(UvSculpt)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(UvSculpt)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "VFont":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*VFont, blk.Hdr.Count)
				for i := range bodies {
					body := new(VFont)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(VFont)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "VPaint":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*VPaint, blk.Hdr.Count)
				for i := range bodies {
					body := new(VPaint)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(VPaint)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "View2D":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*View2D, blk.Hdr.Count)
				for i := range bodies {
					body := new(View2D)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(View2D)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "View3D":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*View3D, blk.Hdr.Count)
				for i := range bodies {
					body := new(View3D)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(View3D)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "VolumeSettings":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*VolumeSettings, blk.Hdr.Count)
				for i := range bodies {
					body := new(VolumeSettings)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(VolumeSettings)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "VoxelData":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*VoxelData, blk.Hdr.Count)
				for i := range bodies {
					body := new(VoxelData)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(VoxelData)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "WarpModifierData":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*WarpModifierData, blk.Hdr.Count)
				for i := range bodies {
					body := new(WarpModifierData)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(WarpModifierData)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "WaveEff":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*WaveEff, blk.Hdr.Count)
				for i := range bodies {
					body := new(WaveEff)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(WaveEff)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "WaveModifierData":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*WaveModifierData, blk.Hdr.Count)
				for i := range bodies {
					body := new(WaveModifierData)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(WaveModifierData)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "WeightVGEditModifierData":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*WeightVGEditModifierData, blk.Hdr.Count)
				for i := range bodies {
					body := new(WeightVGEditModifierData)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(WeightVGEditModifierData)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "WeightVGMixModifierData":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*WeightVGMixModifierData, blk.Hdr.Count)
				for i := range bodies {
					body := new(WeightVGMixModifierData)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(WeightVGMixModifierData)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "WeightVGProximityModifierData":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*WeightVGProximityModifierData, blk.Hdr.Count)
				for i := range bodies {
					body := new(WeightVGProximityModifierData)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(WeightVGProximityModifierData)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "WipeVars":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*WipeVars, blk.Hdr.Count)
				for i := range bodies {
					body := new(WipeVars)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(WipeVars)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "World":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*World, blk.Hdr.Count)
				for i := range bodies {
					body := new(World)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(World)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "bAction":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BAction, blk.Hdr.Count)
				for i := range bodies {
					body := new(BAction)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BAction)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "bActionActuator":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BActionActuator, blk.Hdr.Count)
				for i := range bodies {
					body := new(BActionActuator)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BActionActuator)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "bActionChannel":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BActionChannel, blk.Hdr.Count)
				for i := range bodies {
					body := new(BActionChannel)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BActionChannel)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "bActionConstraint":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BActionConstraint, blk.Hdr.Count)
				for i := range bodies {
					body := new(BActionConstraint)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BActionConstraint)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "bActionGroup":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BActionGroup, blk.Hdr.Count)
				for i := range bodies {
					body := new(BActionGroup)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BActionGroup)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "bActionModifier":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BActionModifier, blk.Hdr.Count)
				for i := range bodies {
					body := new(BActionModifier)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BActionModifier)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "bActionStrip":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BActionStrip, blk.Hdr.Count)
				for i := range bodies {
					body := new(BActionStrip)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BActionStrip)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "bActuator":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BActuator, blk.Hdr.Count)
				for i := range bodies {
					body := new(BActuator)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BActuator)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "bActuatorSensor":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BActuatorSensor, blk.Hdr.Count)
				for i := range bodies {
					body := new(BActuatorSensor)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BActuatorSensor)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "bAddObjectActuator":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BAddObjectActuator, blk.Hdr.Count)
				for i := range bodies {
					body := new(BAddObjectActuator)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BAddObjectActuator)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "bAddon":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BAddon, blk.Hdr.Count)
				for i := range bodies {
					body := new(BAddon)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BAddon)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "bAnimVizSettings":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BAnimVizSettings, blk.Hdr.Count)
				for i := range bodies {
					body := new(BAnimVizSettings)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BAnimVizSettings)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "bArmature":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BArmature, blk.Hdr.Count)
				for i := range bodies {
					body := new(BArmature)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BArmature)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "bArmatureActuator":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BArmatureActuator, blk.Hdr.Count)
				for i := range bodies {
					body := new(BArmatureActuator)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BArmatureActuator)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "bArmatureSensor":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BArmatureSensor, blk.Hdr.Count)
				for i := range bodies {
					body := new(BArmatureSensor)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BArmatureSensor)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "bCameraActuator":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BCameraActuator, blk.Hdr.Count)
				for i := range bodies {
					body := new(BCameraActuator)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BCameraActuator)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "bCameraSolverConstraint":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BCameraSolverConstraint, blk.Hdr.Count)
				for i := range bodies {
					body := new(BCameraSolverConstraint)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BCameraSolverConstraint)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "bChildOfConstraint":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BChildOfConstraint, blk.Hdr.Count)
				for i := range bodies {
					body := new(BChildOfConstraint)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BChildOfConstraint)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "bClampToConstraint":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BClampToConstraint, blk.Hdr.Count)
				for i := range bodies {
					body := new(BClampToConstraint)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BClampToConstraint)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "bCollisionSensor":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BCollisionSensor, blk.Hdr.Count)
				for i := range bodies {
					body := new(BCollisionSensor)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BCollisionSensor)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "bConstraint":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BConstraint, blk.Hdr.Count)
				for i := range bodies {
					body := new(BConstraint)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BConstraint)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "bConstraintActuator":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BConstraintActuator, blk.Hdr.Count)
				for i := range bodies {
					body := new(BConstraintActuator)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BConstraintActuator)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "bConstraintChannel":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BConstraintChannel, blk.Hdr.Count)
				for i := range bodies {
					body := new(BConstraintChannel)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BConstraintChannel)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "bConstraintTarget":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BConstraintTarget, blk.Hdr.Count)
				for i := range bodies {
					body := new(BConstraintTarget)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BConstraintTarget)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "bController":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BController, blk.Hdr.Count)
				for i := range bodies {
					body := new(BController)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BController)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "bDampTrackConstraint":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BDampTrackConstraint, blk.Hdr.Count)
				for i := range bodies {
					body := new(BDampTrackConstraint)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BDampTrackConstraint)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "bDeformGroup":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BDeformGroup, blk.Hdr.Count)
				for i := range bodies {
					body := new(BDeformGroup)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BDeformGroup)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "bDelaySensor":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BDelaySensor, blk.Hdr.Count)
				for i := range bodies {
					body := new(BDelaySensor)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BDelaySensor)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "bDistLimitConstraint":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BDistLimitConstraint, blk.Hdr.Count)
				for i := range bodies {
					body := new(BDistLimitConstraint)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BDistLimitConstraint)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "bDopeSheet":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BDopeSheet, blk.Hdr.Count)
				for i := range bodies {
					body := new(BDopeSheet)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BDopeSheet)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "bEditObjectActuator":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BEditObjectActuator, blk.Hdr.Count)
				for i := range bodies {
					body := new(BEditObjectActuator)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BEditObjectActuator)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "bExpressionCont":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BExpressionCont, blk.Hdr.Count)
				for i := range bodies {
					body := new(BExpressionCont)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BExpressionCont)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "bFollowPathConstraint":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BFollowPathConstraint, blk.Hdr.Count)
				for i := range bodies {
					body := new(BFollowPathConstraint)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BFollowPathConstraint)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "bFollowTrackConstraint":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BFollowTrackConstraint, blk.Hdr.Count)
				for i := range bodies {
					body := new(BFollowTrackConstraint)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BFollowTrackConstraint)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "bGPDframe":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BGPDframe, blk.Hdr.Count)
				for i := range bodies {
					body := new(BGPDframe)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BGPDframe)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "bGPDlayer":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BGPDlayer, blk.Hdr.Count)
				for i := range bodies {
					body := new(BGPDlayer)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BGPDlayer)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "bGPDspoint":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BGPDspoint, blk.Hdr.Count)
				for i := range bodies {
					body := new(BGPDspoint)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BGPDspoint)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "bGPDstroke":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BGPDstroke, blk.Hdr.Count)
				for i := range bodies {
					body := new(BGPDstroke)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BGPDstroke)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "bGPdata":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BGPdata, blk.Hdr.Count)
				for i := range bodies {
					body := new(BGPdata)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BGPdata)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "bGameActuator":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BGameActuator, blk.Hdr.Count)
				for i := range bodies {
					body := new(BGameActuator)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BGameActuator)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "bGroupActuator":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BGroupActuator, blk.Hdr.Count)
				for i := range bodies {
					body := new(BGroupActuator)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BGroupActuator)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "bIKParam":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BIKParam, blk.Hdr.Count)
				for i := range bodies {
					body := new(BIKParam)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BIKParam)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "bIpoActuator":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BIpoActuator, blk.Hdr.Count)
				for i := range bodies {
					body := new(BIpoActuator)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BIpoActuator)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "bItasc":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BItasc, blk.Hdr.Count)
				for i := range bodies {
					body := new(BItasc)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BItasc)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "bJoystickSensor":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BJoystickSensor, blk.Hdr.Count)
				for i := range bodies {
					body := new(BJoystickSensor)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BJoystickSensor)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "bKeyboardSensor":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BKeyboardSensor, blk.Hdr.Count)
				for i := range bodies {
					body := new(BKeyboardSensor)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BKeyboardSensor)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "bKinematicConstraint":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BKinematicConstraint, blk.Hdr.Count)
				for i := range bodies {
					body := new(BKinematicConstraint)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BKinematicConstraint)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "bLocLimitConstraint":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BLocLimitConstraint, blk.Hdr.Count)
				for i := range bodies {
					body := new(BLocLimitConstraint)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BLocLimitConstraint)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "bLocateLikeConstraint":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BLocateLikeConstraint, blk.Hdr.Count)
				for i := range bodies {
					body := new(BLocateLikeConstraint)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BLocateLikeConstraint)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "bLockTrackConstraint":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BLockTrackConstraint, blk.Hdr.Count)
				for i := range bodies {
					body := new(BLockTrackConstraint)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BLockTrackConstraint)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "bMessageActuator":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BMessageActuator, blk.Hdr.Count)
				for i := range bodies {
					body := new(BMessageActuator)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BMessageActuator)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "bMessageSensor":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BMessageSensor, blk.Hdr.Count)
				for i := range bodies {
					body := new(BMessageSensor)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BMessageSensor)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "bMinMaxConstraint":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BMinMaxConstraint, blk.Hdr.Count)
				for i := range bodies {
					body := new(BMinMaxConstraint)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BMinMaxConstraint)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "bMotionPath":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BMotionPath, blk.Hdr.Count)
				for i := range bodies {
					body := new(BMotionPath)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BMotionPath)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "bMotionPathVert":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BMotionPathVert, blk.Hdr.Count)
				for i := range bodies {
					body := new(BMotionPathVert)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BMotionPathVert)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "bMouseSensor":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BMouseSensor, blk.Hdr.Count)
				for i := range bodies {
					body := new(BMouseSensor)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BMouseSensor)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "bNearSensor":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BNearSensor, blk.Hdr.Count)
				for i := range bodies {
					body := new(BNearSensor)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BNearSensor)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "bNode":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BNode, blk.Hdr.Count)
				for i := range bodies {
					body := new(BNode)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BNode)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "bNodeInstanceHashEntry":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BNodeInstanceHashEntry, blk.Hdr.Count)
				for i := range bodies {
					body := new(BNodeInstanceHashEntry)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BNodeInstanceHashEntry)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "bNodeInstanceKey":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BNodeInstanceKey, blk.Hdr.Count)
				for i := range bodies {
					body := new(BNodeInstanceKey)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BNodeInstanceKey)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "bNodeLink":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BNodeLink, blk.Hdr.Count)
				for i := range bodies {
					body := new(BNodeLink)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BNodeLink)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "bNodePreview":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BNodePreview, blk.Hdr.Count)
				for i := range bodies {
					body := new(BNodePreview)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BNodePreview)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "bNodeSocket":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BNodeSocket, blk.Hdr.Count)
				for i := range bodies {
					body := new(BNodeSocket)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BNodeSocket)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "bNodeSocketValueBoolean":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BNodeSocketValueBoolean, blk.Hdr.Count)
				for i := range bodies {
					body := new(BNodeSocketValueBoolean)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BNodeSocketValueBoolean)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "bNodeSocketValueFloat":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BNodeSocketValueFloat, blk.Hdr.Count)
				for i := range bodies {
					body := new(BNodeSocketValueFloat)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BNodeSocketValueFloat)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "bNodeSocketValueInt":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BNodeSocketValueInt, blk.Hdr.Count)
				for i := range bodies {
					body := new(BNodeSocketValueInt)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BNodeSocketValueInt)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "bNodeSocketValueRGBA":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BNodeSocketValueRGBA, blk.Hdr.Count)
				for i := range bodies {
					body := new(BNodeSocketValueRGBA)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BNodeSocketValueRGBA)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "bNodeSocketValueString":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BNodeSocketValueString, blk.Hdr.Count)
				for i := range bodies {
					body := new(BNodeSocketValueString)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BNodeSocketValueString)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "bNodeSocketValueVector":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BNodeSocketValueVector, blk.Hdr.Count)
				for i := range bodies {
					body := new(BNodeSocketValueVector)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BNodeSocketValueVector)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "bNodeStack":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BNodeStack, blk.Hdr.Count)
				for i := range bodies {
					body := new(BNodeStack)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BNodeStack)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "bNodeTree":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BNodeTree, blk.Hdr.Count)
				for i := range bodies {
					body := new(BNodeTree)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BNodeTree)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "bNodeTreePath":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BNodeTreePath, blk.Hdr.Count)
				for i := range bodies {
					body := new(BNodeTreePath)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BNodeTreePath)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "bObjectActuator":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BObjectActuator, blk.Hdr.Count)
				for i := range bodies {
					body := new(BObjectActuator)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BObjectActuator)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "bObjectSolverConstraint":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BObjectSolverConstraint, blk.Hdr.Count)
				for i := range bodies {
					body := new(BObjectSolverConstraint)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BObjectSolverConstraint)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "bParentActuator":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BParentActuator, blk.Hdr.Count)
				for i := range bodies {
					body := new(BParentActuator)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BParentActuator)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "bPivotConstraint":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BPivotConstraint, blk.Hdr.Count)
				for i := range bodies {
					body := new(BPivotConstraint)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BPivotConstraint)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "bPose":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BPose, blk.Hdr.Count)
				for i := range bodies {
					body := new(BPose)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BPose)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "bPoseChannel":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BPoseChannel, blk.Hdr.Count)
				for i := range bodies {
					body := new(BPoseChannel)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BPoseChannel)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "bProperty":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BProperty, blk.Hdr.Count)
				for i := range bodies {
					body := new(BProperty)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BProperty)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "bPropertyActuator":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BPropertyActuator, blk.Hdr.Count)
				for i := range bodies {
					body := new(BPropertyActuator)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BPropertyActuator)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "bPropertySensor":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BPropertySensor, blk.Hdr.Count)
				for i := range bodies {
					body := new(BPropertySensor)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BPropertySensor)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "bPythonConstraint":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BPythonConstraint, blk.Hdr.Count)
				for i := range bodies {
					body := new(BPythonConstraint)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BPythonConstraint)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "bPythonCont":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BPythonCont, blk.Hdr.Count)
				for i := range bodies {
					body := new(BPythonCont)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BPythonCont)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "bRadarSensor":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BRadarSensor, blk.Hdr.Count)
				for i := range bodies {
					body := new(BRadarSensor)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BRadarSensor)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "bRandomActuator":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BRandomActuator, blk.Hdr.Count)
				for i := range bodies {
					body := new(BRandomActuator)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BRandomActuator)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "bRandomSensor":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BRandomSensor, blk.Hdr.Count)
				for i := range bodies {
					body := new(BRandomSensor)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BRandomSensor)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "bRaySensor":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BRaySensor, blk.Hdr.Count)
				for i := range bodies {
					body := new(BRaySensor)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BRaySensor)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "bRigidBodyJointConstraint":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BRigidBodyJointConstraint, blk.Hdr.Count)
				for i := range bodies {
					body := new(BRigidBodyJointConstraint)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BRigidBodyJointConstraint)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "bRotLimitConstraint":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BRotLimitConstraint, blk.Hdr.Count)
				for i := range bodies {
					body := new(BRotLimitConstraint)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BRotLimitConstraint)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "bRotateLikeConstraint":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BRotateLikeConstraint, blk.Hdr.Count)
				for i := range bodies {
					body := new(BRotateLikeConstraint)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BRotateLikeConstraint)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "bSameVolumeConstraint":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BSameVolumeConstraint, blk.Hdr.Count)
				for i := range bodies {
					body := new(BSameVolumeConstraint)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BSameVolumeConstraint)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "bSceneActuator":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BSceneActuator, blk.Hdr.Count)
				for i := range bodies {
					body := new(BSceneActuator)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BSceneActuator)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "bScreen":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BScreen, blk.Hdr.Count)
				for i := range bodies {
					body := new(BScreen)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BScreen)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "bSensor":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BSensor, blk.Hdr.Count)
				for i := range bodies {
					body := new(BSensor)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BSensor)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "bShrinkwrapConstraint":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BShrinkwrapConstraint, blk.Hdr.Count)
				for i := range bodies {
					body := new(BShrinkwrapConstraint)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BShrinkwrapConstraint)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "bSizeLikeConstraint":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BSizeLikeConstraint, blk.Hdr.Count)
				for i := range bodies {
					body := new(BSizeLikeConstraint)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BSizeLikeConstraint)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "bSizeLimitConstraint":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BSizeLimitConstraint, blk.Hdr.Count)
				for i := range bodies {
					body := new(BSizeLimitConstraint)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BSizeLimitConstraint)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "bSound":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BSound, blk.Hdr.Count)
				for i := range bodies {
					body := new(BSound)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BSound)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "bSoundActuator":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BSoundActuator, blk.Hdr.Count)
				for i := range bodies {
					body := new(BSoundActuator)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BSoundActuator)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "bSplineIKConstraint":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BSplineIKConstraint, blk.Hdr.Count)
				for i := range bodies {
					body := new(BSplineIKConstraint)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BSplineIKConstraint)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "bStateActuator":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BStateActuator, blk.Hdr.Count)
				for i := range bodies {
					body := new(BStateActuator)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BStateActuator)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "bStats":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BStats, blk.Hdr.Count)
				for i := range bodies {
					body := new(BStats)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BStats)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "bSteeringActuator":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BSteeringActuator, blk.Hdr.Count)
				for i := range bodies {
					body := new(BSteeringActuator)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BSteeringActuator)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "bStretchToConstraint":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BStretchToConstraint, blk.Hdr.Count)
				for i := range bodies {
					body := new(BStretchToConstraint)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BStretchToConstraint)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "bTheme":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BTheme, blk.Hdr.Count)
				for i := range bodies {
					body := new(BTheme)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BTheme)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "bTouchSensor":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BTouchSensor, blk.Hdr.Count)
				for i := range bodies {
					body := new(BTouchSensor)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BTouchSensor)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "bTrackToConstraint":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BTrackToConstraint, blk.Hdr.Count)
				for i := range bodies {
					body := new(BTrackToConstraint)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BTrackToConstraint)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "bTransLikeConstraint":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BTransLikeConstraint, blk.Hdr.Count)
				for i := range bodies {
					body := new(BTransLikeConstraint)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BTransLikeConstraint)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "bTransformConstraint":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BTransformConstraint, blk.Hdr.Count)
				for i := range bodies {
					body := new(BTransformConstraint)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BTransformConstraint)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "bTwoDFilterActuator":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BTwoDFilterActuator, blk.Hdr.Count)
				for i := range bodies {
					body := new(BTwoDFilterActuator)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BTwoDFilterActuator)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "bVisibilityActuator":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*BVisibilityActuator, blk.Hdr.Count)
				for i := range bodies {
					body := new(BVisibilityActuator)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(BVisibilityActuator)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "rctf":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*Rctf, blk.Hdr.Count)
				for i := range bodies {
					body := new(Rctf)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(Rctf)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "rcti":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*Rcti, blk.Hdr.Count)
				for i := range bodies {
					body := new(Rcti)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(Rcti)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "uiFont":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*UiFont, blk.Hdr.Count)
				for i := range bodies {
					body := new(UiFont)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(UiFont)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "uiFontStyle":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*UiFontStyle, blk.Hdr.Count)
				for i := range bodies {
					body := new(UiFontStyle)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(UiFontStyle)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "uiGradientColors":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*UiGradientColors, blk.Hdr.Count)
				for i := range bodies {
					body := new(UiGradientColors)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(UiGradientColors)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "uiList":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*UiList, blk.Hdr.Count)
				for i := range bodies {
					body := new(UiList)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(UiList)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "uiPanelColors":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*UiPanelColors, blk.Hdr.Count)
				for i := range bodies {
					body := new(UiPanelColors)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(UiPanelColors)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "uiStyle":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*UiStyle, blk.Hdr.Count)
				for i := range bodies {
					body := new(UiStyle)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(UiStyle)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "uiWidgetColors":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*UiWidgetColors, blk.Hdr.Count)
				for i := range bodies {
					body := new(UiWidgetColors)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(UiWidgetColors)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "uiWidgetStateColors":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*UiWidgetStateColors, blk.Hdr.Count)
				for i := range bodies {
					body := new(UiWidgetStateColors)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(UiWidgetStateColors)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "vec2f":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*Vec2f, blk.Hdr.Count)
				for i := range bodies {
					body := new(Vec2f)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(Vec2f)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "vec2s":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*Vec2s, blk.Hdr.Count)
				for i := range bodies {
					body := new(Vec2s)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(Vec2s)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "vec3f":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*Vec3f, blk.Hdr.Count)
				for i := range bodies {
					body := new(Vec3f)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(Vec3f)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "wmKeyConfig":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*WmKeyConfig, blk.Hdr.Count)
				for i := range bodies {
					body := new(WmKeyConfig)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(WmKeyConfig)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "wmKeyMap":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*WmKeyMap, blk.Hdr.Count)
				for i := range bodies {
					body := new(WmKeyMap)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(WmKeyMap)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "wmKeyMapDiffItem":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*WmKeyMapDiffItem, blk.Hdr.Count)
				for i := range bodies {
					body := new(WmKeyMapDiffItem)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(WmKeyMapDiffItem)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "wmKeyMapItem":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*WmKeyMapItem, blk.Hdr.Count)
				for i := range bodies {
					body := new(WmKeyMapItem)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(WmKeyMapItem)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "wmOperator":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*WmOperator, blk.Hdr.Count)
				for i := range bodies {
					body := new(WmOperator)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(WmOperator)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "wmWindow":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*WmWindow, blk.Hdr.Count)
				for i := range bodies {
					body := new(WmWindow)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(WmWindow)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		case "wmWindowManager":
			if blk.Hdr.Count > 1 {
				// Parse block body structures.
				bodies := make([]*WmWindowManager, blk.Hdr.Count)
				for i := range bodies {
					body := new(WmWindowManager)
					err = binary.Read(r, order, body)
					if err != nil {
						return err
					}
					bodies[i] = body
				}
				blk.Body = bodies
			} else {
				// Parse block body structure.
				body := new(WmWindowManager)
				err = binary.Read(r, order, body)
				if err != nil {
					return err
				}
				blk.Body = body
			}
			/// ### [ tmp ] ###
			// Verify that all bytes in the block body have been read.
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			if len(buf) > 0 {
				log.Printf("%d unread bytes in %q.", len(buf), typ)
				log.Printf("blk.Hdr: %#v\n", blk.Hdr)
				log.Println(hex.Dump(buf))
				os.Exit(1)
			}
			/// ### [/ tmp ] ###
		}
	}

	return nil
}
